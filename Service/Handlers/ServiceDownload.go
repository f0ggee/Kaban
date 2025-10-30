package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type CustomError struct {
	Message string
	Err     error
}

func getNameFromUrl(r *http.Request) string {
	vars := mux.Vars(r)

	name := vars["name"]
	slog.Info("File Name", name)
	return name

}

func (d *CustomError) ServiceDownload(w http.ResponseWriter, r *http.Request, ch chan string) *CustomError {

	ctx, cancle := Uttiltesss.Contexte()
	defer cancle()

	name := getNameFromUrl(r)
	if name == " " {
		return &CustomError{
			Message: "Error name don't set\"",
			Err:     nil,
		}
	}
	bucket := os.Getenv("BUCKET")

	sees, err := Uttiltesss.Inzelire()
	if err != nil {
		slog.Error("Error in func", err)
		return &CustomError{
			Message: "Error in set S3",
			Err:     err,
		}
	}

	downloader := s3.New(sees)

	o, err := downloader.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket:      aws.String(bucket),
		IfNoneMatch: aws.String(""),
		Key:         aws.String(name),
	})
	defer o.Body.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename= %v", name))

	if _, err = io.Copy(w, o.Body); err != nil {
		slog.Error("Err In file Service Downloader", err)
		return &CustomError{
			Message: "Err in downloadin file",
			Err:     err,
		}
	}
	ch <- name
	return nil
}
