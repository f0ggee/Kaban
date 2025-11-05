package Controller

import (
	"Kaban/Service/Handlers"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"time"
)

func ControlerFileUploader(w http.ResponseWriter, r *http.Request, router *mux.Router) {
	if r.Method != http.MethodPost {
		http.Error(w, "err", http.StatusUnauthorized)
		slog.Error("Err in Cottroler Uploader")
		return
	}

	timeR := time.Now()
	filName, err := Handlers.FileUploader(w, r)
	if err != nil {
		slog.Error("Error in file uploader", err)
		return
	}

	url, err := router.Get("fileName").URL("name", filName)
	if err != nil {
		slog.Error("Error can't treate", err)
		return
	}
	e := time.Since(timeR)
	slog.Info("Work of function - ", "time:", e)
	http.Redirect(w, r, url.Path, http.StatusFound)

}
