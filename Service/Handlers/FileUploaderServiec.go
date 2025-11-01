package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"bytes"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/gorilla/mux"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
)

func FileUploader(w http.ResponseWriter, r *http.Request, router *mux.Router) error {

	cfg, err := Uttiltesss.Inzelire2()
	if err != nil {
		slog.Error("Err cant", err)
		return err
	}

	ctx, cancel := Uttiltesss.Contexte()
	defer cancel()

	sizeAndName, f, err2 := checkAndDownload(w, r, err)
	if err2 != nil {
		slog.Error("Error in fileUploader", err2)
		return err2
	}
	bucket := "0c8f1ea9-b07f5996-b392-4227-961b-14d2a71a53dc"
	uploade := manager.NewUploader(cfg, func(uploader *manager.Uploader) {
		uploader.MaxUploadParts = 100 * 1024 * 1024
		uploader.PartSize = 50 * 1024 * 1024
		uploader.Concurrency = 3

	})

	_, err = uploade.Upload(ctx, &s3.PutObjectInput{Bucket: aws.String(bucket), Key: aws.String(sizeAndName.Filename), Body: bytes.NewReader(f)})
	if err != nil {
		slog.Error("Error", err)
	}

	slog.Info("File succes upload :)")

	url, err := router.Get("fileName").URL("name", sizeAndName.Filename)
	if err != nil {
		slog.Error("Erro can't treate", err)
		return err
	}
	http.Redirect(w, r, url.Path, http.StatusFound)

	return nil

}

func checkAndDownload(w http.ResponseWriter, r *http.Request, err error) (*multipart.FileHeader, []byte, error) {
	file, sizeAndName, err := r.FormFile("file")
	if err != nil {
		slog.Error("Err from FileUploader 1 ", err)
		http.Error(w, "Error", http.StatusNotFound)
		return nil, nil, err
	}
	f, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Error in file Uploader 3 ")
		return nil, nil, err
	}
	if sizeAndName.Size > 5000000000 {
		http.Error(w, "File can't be treate", http.StatusUnauthorized)
		slog.Error("Error in file it's too big  Uploader 2 ", err)
		return nil, nil, err
	}

	defer file.Close()
	return sizeAndName, f, nil
}
