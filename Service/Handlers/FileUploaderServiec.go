package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"bufio"
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

var n []byte
var Nonce = make(map[string]string)

func FileUploader(w http.ResponseWriter, r *http.Request, sc string) (string, error) {

	reader, writer := io.Pipe()

	ch := make(chan []byte)

	file, sizeAndName, err := r.FormFile("file")
	if err != nil {
		slog.Error("Err from FileUploader 1 ", err)
		http.Error(w, "Error", http.StatusNotFound)
		return "", err
	}
	defer file.Close()
	if sizeAndName.Size > 5000000000 {
		http.Error(w, "File can't be treate", http.StatusUnauthorized)
		slog.Error("Error in file it's too big  Uploader 2 ", err)
		return "", err
	}

	cfg, err := Uttiltesss.Inzelire2()
	if err != nil {
		slog.Error("Err cant", err)
		return "", err
	}

	go func() {
		defer writer.Close()
		err = Encrypt(ch, sc, file, writer)
		if err != nil {
			writer.CloseWithError(err)

		}

	}()

	ctx, cancel := Uttiltesss.Contexte()
	defer cancel()

	bucket := "0c8f1ea9-b07f5996-b392-4227-961b-14d2a71a53dc"
	uploade := manager.NewUploader(cfg, func(uploader *manager.Uploader) {
		uploader.MaxUploadParts = 1000
		uploader.PartSize = 10 * 1024 * 1024
		uploader.Concurrency = 25
	})

	_, err = uploade.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(sizeAndName.Filename),
		Body:   reader,
	})
	if err != nil {
		slog.Error("Error in uploader", err)
		return "", err
	}

	nonce := <-ch
	n = nonce
	//Nonce[sizeAndName.Filename] = hex.EncodeToString(nonce)

	slog.Info("File succes upload :)")

	return sizeAndName.Filename, nil

}

func Encrypt(ch chan []byte, SC string, file multipart.File, writer io.Writer) error {
	Sc, err := hex.DecodeString(SC)
	if err != nil {
		slog.Error("Error Decode string", err)
		return err
	}
	block, err := aes.NewCipher(Sc)
	if err != nil {
		slog.Error("Error in Create New Block", err)
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		slog.Error("Error in create GCM mode", err)
		return err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		slog.Error("Error in creating nonce", err)
		return err
	}

	p := make([]byte, 10)
	ze := bufio.NewReader(file)

	for {
		n, err := ze.Read(p)
		if n > 0 {
			cipherT := gcm.Seal(nil, nonce, p, nil)
			_, errW := writer.Write(cipherT)
			if errW != nil {
				slog.Error("Error in 108", err)
				return errW
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("Error", "error ", err)
			return err
		}

	}

	go func() {
		ch <- nonce
	}()
	return nil
}
