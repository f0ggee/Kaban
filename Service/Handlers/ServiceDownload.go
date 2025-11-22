package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"os"
)

type CustomError struct {
	Message string
	Err     error
}

func ServiceDownload(w http.ResponseWriter, r *http.Request, sc string, name string) error {

	ctx, cancle := Uttiltesss.Contexte()
	defer cancle()

	//
	SC, ok := Nonce[name]
	if !ok {
		return errors.New("Nonce is don't set ")
	}

	Reader, writer := io.Pipe()

	bucket := os.Getenv("BUCKET")

	sees, err := Uttiltesss.Inzelire()
	if err != nil {
		slog.Error("Error in func", err)

		return err
	}

	downloader := s3.New(sees)

	o, err := downloader.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket:      aws.String(bucket),
		IfNoneMatch: aws.String(""),
		Key:         &name,
	})

	if err != nil {
		slog.Error("ServiceDownload:", err)
		return err
	}

	go func() {
		defer writer.Close()
		err = DecryptFile(SC, o, writer)
		if err != nil {
			err := writer.CloseWithError(err)
			if err != nil {
				slog.Error("Error", "err", err)
				return
			}

		}
	}()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename= %v", name))
	w.Header().Set("Content-Length", strconv.FormatInt(*o.ContentLength-aes.BlockSize, 10))

	if _, err = io.Copy(w, Reader); err != nil {
		slog.Error("Err In file Service Downloader", err)
		return err
	}

	return nil
}

func DecryptFile(SC string, o *s3.GetObjectOutput, writer *io.PipeWriter) error {

	sc, err := hex.DecodeString(SC)
	if err != nil {
		slog.Error("Error in decrypt file ", err)
		return err
	}
	block, err := aes.NewCipher(sc)
	if err != nil {
		slog.Error("Error in  create file", err)
		return err
	}

	nonce := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(o.Body, nonce)
	if err != nil {
		slog.Error("Error in read", err)
		return err
	}

	plaintext := make([]byte, 35*1024)

	stream := cipher.NewCTR(block, nonce)

	file := bufio.NewReader(o.Body)
	for {
		n, err := file.Read(plaintext)
		if err != nil && err != io.EOF {
			slog.Error("Error in file", err)
			return err
		}
		if err == io.EOF {
			break
		}

		if n > 0 {
			stream.XORKeyStream(plaintext[:n], plaintext[:n])
			_, err = writer.Write(plaintext[:n])
			if err != nil {
				writer.CloseWithError(err)
				return err
			}
		}

	}

	return nil
}
