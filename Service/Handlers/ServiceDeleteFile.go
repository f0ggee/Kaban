package Handlers

import (
	"Kaban/Service/Uttiltesss"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log/slog"
	"os"
)

func (d *CustomError) Delete(ch chan string) *CustomError {

	nameOfFile := <-ch
	bucket := os.Getenv("BUCKET")
	if nameOfFile == "" {
		return &CustomError{
			Message: "Name Of file don't exits",
			Err:     nil,
		}
	}

	_ = s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(nameOfFile)}
}
