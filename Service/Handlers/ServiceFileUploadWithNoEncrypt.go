package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"log/slog"
	"mime/multipart"
	"net/http"
)

func FileUploaderNoEncr(w http.ResponseWriter, r *http.Request) (string, error) {
	ctx, cancel := Uttiltesss.Contexte()
	defer cancel()

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

	err = CheckFileSize2(w, sizeAndName, err)
	if err != nil {
		slog.Error("Error in Check File size", err)
		return "", err
	}

	cfg, err := Uttiltesss.Inzelire2()
	if err != nil {
		slog.Error("Err cant", err)
		return "", err
	}

	bucket := "0c8f1ea9-b07f5996-b392-4227-961b-14d2a71a53dc"
	uploader := manager.NewUploader(cfg, func(uploader *manager.Uploader) {
		uploader.MaxUploadParts = 1000
		uploader.PartSize = 10 * 1024 * 1024
		uploader.Concurrency = 25
	})

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(sizeAndName.Filename),
		Body:   file,
	})
	if err != nil {
		slog.Error("Error in uploader", err)
		return "", err
	}

	slog.Info("File success upload :)")

	return sizeAndName.Filename, nil

}

func CheckFileSize2(w http.ResponseWriter, sizeAndName *multipart.FileHeader, err error) error {
	if sizeAndName.Size > 2000000000 {
		http.Error(w, "File can't be treating", http.StatusUnauthorized)
		slog.Error("Error in file it's too big  Uploader 2 ", err)
		return err
	}
	return nil
}
