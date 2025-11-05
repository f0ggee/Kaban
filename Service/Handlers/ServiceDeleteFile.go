package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"

	"log/slog"
	"os"
)

func Delete(ch chan string) {

	slog.Info("I'm here ")
	name := <-ch
	s, err := Uttiltesss.Inzelire2()
	if err != nil {
		slog.Error("Eror in func Delete file", "err", err)
		return
	}
	ctx, canclet := Uttiltesss.Contexte()
	defer canclet()

	params := s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(name),
	}
	sz, err := s.DeleteObject(ctx, &params)
	if err != nil {
		slog.Error("Error in Delete file 2", err)
		return
	}

	slog.Info("Result", "res", sz.ResultMetadata)

}
