package Controller

import (
	"Kaban/Service/Handlers"
	"errors"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

func ControlerFileUploaderNoEncrypt(w http.ResponseWriter, r *http.Request, router *mux.Router) {
	if r.Method != http.MethodPost {
		http.Error(w, "err", http.StatusUnauthorized)
		slog.Error("Err in Cottroler Uploader")
		return
	}

	err := CookiGet2(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)

	}

	filName, err := Handlers.FileUploaderNoEncr(w, r)
	if err != nil {
		return
	}

	url, err := router.Get("fileName").URL("name", filName, "bool", "false")
	if err != nil {
		slog.Error("Error can't treate", err)
		return
	}

	slog.Info(url.Path)
	http.Redirect(w, r, url.Path, http.StatusFound)

}

func CookiGet2(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		http.Error(w, "cookie dont sen", http.StatusUnauthorized)
		return err
	}

	if session.Options.MaxAge == 0 {
		slog.Error("Cookie time expired")
		return errors.New("Cookie time expired")
	}

	return nil
}
