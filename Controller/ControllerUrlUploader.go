package Controller

import (
	"Kaban/Service/Handlers"
	"log/slog"
	"net/http"
)

func CUrlUp(w http.ResponseWriter, r *http.Request) string {
	if r.Method != http.MethodGet {
		slog.Error("Method don't allow")
		http.Error(w, "Method", http.StatusUnauthorized)
		return ""
	}

	nameFile := Handlers.UrlUploader(r)
	if nameFile == "" {
		return ""
	}
	return nameFile

}
