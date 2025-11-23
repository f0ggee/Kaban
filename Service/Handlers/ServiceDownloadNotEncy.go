package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"crypto/aes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"os"
)

func SDownloadWithNonEncrypt(w http.ResponseWriter, name string) error {

	ctx, cancle := Uttiltesss.Contexte()
	defer cancle()

	bucket := os.Getenv("BUCKET")

	sees, err := Uttiltesss.Inzelire()
	if err != nil {
		slog.Error("Eror in inzelizre s3 server", "err:", err)

		return err
	}

	downloader := s3.New(sees)

	o, err := downloader.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket:      aws.String(bucket),
		IfNoneMatch: aws.String(""),
		Key:         &name,
	})

	if err != nil {
		slog.Error("ServiceDownload:", "err", err)
		return err
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename= %v", name))
	w.Header().Set("Content-Length", strconv.FormatInt(*o.ContentLength-aes.BlockSize, 10))

	if _, err = io.Copy(w, o.Body); err != nil {
		slog.Error("Err In file Service Downloader", "err", err)
		return err
	}

	return nil
}
