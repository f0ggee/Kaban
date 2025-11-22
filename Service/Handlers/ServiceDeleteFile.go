package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"

	"log/slog"
	"os"
)

func Delete(nameOfKey string) {

	slog.Info("I'm here ")
	s, err := Uttiltesss.Inzelire2()
	if err != nil {
		slog.Error("Eror in func Delete file", "err", err)
		return
	}
	ctx, canclet := Uttiltesss.Contexte()
	defer canclet()

	params := s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(nameOfKey),
	}
	_, err = s.DeleteObject(ctx, &params)
	if err != nil {
		slog.Error("Error in Delete file 2", err)
		return
	}

	if _, ok := Nonce[nameOfKey]; ok {
		delete(Nonce, nameOfKey)

	} else {
		slog.Info("Result", "Delete", false)
	}

	slog.Info("Result", "Delete", true)
}
