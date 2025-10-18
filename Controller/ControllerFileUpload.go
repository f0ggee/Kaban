package Controller

import (
	"Kaban/Service/Handlers"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

func ControlerFileUploader(w http.ResponseWriter, r *http.Request, router *mux.Router) {
	if r.Method != http.MethodPost {
		http.Error(w, "err", http.StatusUnauthorized)
		slog.Error("Err in Cottroler Uploader")
		return
	}

	err := Handlers.FileUploader(w, r, router)
	if err != nil {
		slog.Error("Error in file uploader", err)
		return
	}

}
