package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
)

func FileUploaderEnc(w http.ResponseWriter, r *http.Request, sc string) (string, error) {
	ctx, cancel := Uttiltesss.Contexte()
	defer cancel()

	reader, writer := io.Pipe()

	file, sizeAndName, err := r.FormFile("file")
	if err != nil {
		slog.Error("Err from FileUploader 1 ", err)
		http.Error(w, "Error", http.StatusNotFound)
		return "", err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			slog.Error("Err, cant' close a file", "err", err)
			return
		}
	}()

	err = CheckFileSize(w, sizeAndName, err)
	if err != nil {
		slog.Error("Error in Check File size", err)
		return "", err
	}

	cfg, err := Uttiltesss.Inzelire2()
	if err != nil {
		slog.Error("Err cant", err)
		return "", err
	}

	go func() {
		defer writer.Close()
		err = Encrypt(sc, file, writer)
		if err != nil {
			err := writer.CloseWithError(err)
			if err != nil {
				slog.Error("Error in file writing", err)
				return
			}

		}
	}()

	bucket := "0c8f1ea9-b07f5996-b392-4227-961b-14d2a71a53dc"
	uploader := manager.NewUploader(cfg, func(uploader *manager.Uploader) {
		uploader.MaxUploadParts = 1000
		uploader.PartSize = 10 * 1024 * 1024
		uploader.Concurrency = 25
	})

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(sizeAndName.Filename),
		Body:   reader,
	})
	if err != nil {
		slog.Error("Error in uploader", err)
		return "", err
	}

	slog.Info("File success upload :)")

	go func() {
		Nonce = make(map[string]string)
		Nonce[sizeAndName.Filename] = sc
	}()

	return sizeAndName.Filename, nil

}

func CheckFileSize(w http.ResponseWriter, sizeAndName *multipart.FileHeader, err error) error {
	if sizeAndName.Size > 2000000000 {
		http.Error(w, "File can't be treating", http.StatusUnauthorized)
		slog.Error("Error in file it's too big  Uploader 2 ", err)
		return err
	}
	return nil
}

func Encrypt(SC string, file multipart.File, writer io.Writer) error {
	Sc, err := hex.DecodeString(SC)
	if err != nil {
		slog.Error("Error Decode string", err)
		return err
	}
	slog.Info("key", SC)

	block, err := aes.NewCipher(Sc)
	if err != nil {
		slog.Error("err create a NewCipher")
		return err
	}

	nonce := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, nonce)
	buf := make([]byte, 32*1024)
	writer.Write(nonce)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			slog.Error("Error in file upload", err)
			return err
		}
		if err == io.EOF {
			break
		}
		stream.XORKeyStream(buf[:n], buf[:n])
		writer.Write(buf[:n])

	}

	return nil
}
