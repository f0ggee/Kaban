package Controller

import (
	"Kaban/Service/Handlers"
	"errors"
	"fmt"
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

	SC, err := CookiGet(w, r)
	if err != nil {
		http.Error(w, "Cookie don't find", http.StatusNotFound)

	}

	filName, err := Handlers.FileUploader(w, r, SC)
	if err != nil {
		return
	}

	url, err := router.Get("fileName").URL("name", filName)
	if err != nil {
		slog.Error("Error can't treate", err)
		return
	}

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

	_, ok := session.Values["cookie"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)

	}
	SC, ok := session.Values["SW"]
	if !ok {
		slog.Error("Err", "EXIST:", ok)
		return "", errors.New("SC don't set")
	}

	scs := fmt.Sprint(SC)
	return scs, err
}
