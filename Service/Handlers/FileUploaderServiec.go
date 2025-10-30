package Handlers

import (
	"Kaban/Service/Connect_to_BD"
	"Kaban/Service/Uttiltesss"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"io"
	"log/slog"
	"net/http"
)

//func dropInTheUrl(fileName string, router *mux.Router) error {
//
//	_, err := router.Get("UserFileName").URL("name", fileName)
//	if err != nil {
//		slog.Error("Error in dropInTheUrl ", err)
//		return err
//	}
//
//	return nil
//}

func getCookie(r *http.Request) (string, error) {
	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		return "", err

	}

	cookie, ok := session.Values["cookie"]
	if !ok {

		slog.Error("Cookie don't set", err)
		return "", err
	}
	s := fmt.Sprintf("string", cookie)

	return s, nil
}

func getKey(r *http.Request) (string, error) {

	db, err := Connect_to_BD.Connect()
	if err != nil {
		slog.Error("Err in GetKey 1", err)
		return "", err
	}

	s, err2 := getCookie(r)
	if err2 != nil {
		slog.Error("Error in GetKey 2 ", err)
	}
	var (
		sctypt string
	)

	err = db.QueryRow(context.Background(), `SELECT  scrypt_salt FROM person WHERE cookie=$1`, s).Scan(&sctypt)
	if err != nil {
		slog.Error("Error", err)
		return "", err
	}

	return sctypt, nil
}

func FileUploader(w http.ResponseWriter, r *http.Request, router *mux.Router) error {

	s3, err := Uttiltesss.Inzelire()
	if err != nil {
		slog.Info("Can't set connection in s3 server", err)
		return err
	}

	//hexString, err := getKey(r)
	//if err != nil {
	//	slog.Error("Error can't get scrypt from Database", err)
	//	return
	//}
	file, sizeAndName, err := r.FormFile("file")
	if err != nil {
		slog.Error("Err from FileUploader 1 ", err)
		http.Error(w, "Error", http.StatusNotFound)
		return err
	}
	f, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Error in file Uploader 3 ")
		return err
	}
	if sizeAndName.Size > 5000000000 {
		http.Error(w, "File can't be treate", http.StatusUnauthorized)
		slog.Error("Error in file it's too big  Uploader 2 ", err)
		return err
	}

	defer file.Close()

	uploader := s3manager.NewUploader(s3)

	bucket := "0c8f1ea9-b07f5996-b392-4227-961b-14d2a71a53dc"
	up := &s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &sizeAndName.Filename,
		Body:   bytes.NewReader(f),
	}
	_, err = uploader.Upload(up)
	if err != nil {
		return err
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

func Sha256(f []byte) string {
	hash := sha256.Sum256(f)
	h := hex.EncodeToString(hash[:])
	return h
}
