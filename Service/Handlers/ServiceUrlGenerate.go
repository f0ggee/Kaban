package Handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func UrlUploader(r *http.Request) string {
	name := r.URL.Query().Get("name")
	fmt.Println(Nonce[name])

	slog.Info("From UrlUploader  handler", "Filename - ", name)
	return name
}
