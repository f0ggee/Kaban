package Handlers

import (
	"log/slog"
	"net/http"
)

func UrlUploader(r *http.Request) string {
	name := r.URL.Query().Get("name")

	slog.Info("From UrlUploader  handler :File name ", name)
	return name
}
