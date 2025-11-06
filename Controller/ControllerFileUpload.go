package Controller

import (
	"Kaban/Service/Handlers"
	"errors"
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

	SC, err := CookiGet(w, r)
	if err != nil {
		slog.Error("Error cant ControlerFileUploader", err)
		return
	}

	timeR := time.Now()
	filName, err := Handlers.FileUploader(w, r, SC)
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
	slog.Info(url.Path)
	http.Redirect(w, r, url.Path, http.StatusFound)

}

func CookiGet(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		http.Error(w, "cookie dont sen", http.StatusUnauthorized)
		return "", err
	}
	if session.Options.MaxAge == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	SC, ok := session.Values["SC"].(string)
	if !ok {
		slog.Error("Err", "EXIST:", ok)
		return "", errors.New("SC don't set")
	}
	return SC, err
}
