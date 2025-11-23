package Controller

import (
	"Kaban/Service/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

func CookiGetInControllerDownloader(w http.ResponseWriter, r *http.Request) (string, error) {
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
	SC, _ := session.Values["SW"].(string)

	slog.Info("SC")
	return SC, err
}
func getNameFromUrl(r *http.Request) string {
	vars := mux.Vars(r)

	name := vars["name"]
	slog.Info(name)
	return name

}
func DownloadWithEncrypt(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Status method don't allow", http.StatusBadRequest)
		return
	}

	name := getNameFromUrl(r)

	sc, er := CookiGetInControllerDownloader(w, r)
	if er != nil {
		slog.Error("Can't get cookie because", "err", er)
		return
	}

	err := Handlers.SDownloadEncrypt(w, r, sc, name)
	if err != nil {
		_, err = fmt.Fprintf(w, "Can't dowload file")
		http.Error(w, "Error because"+fmt.Sprint(err), http.StatusBadRequest)
		if err != nil {
			slog.Error("Error writing a message", err)
			return
		}
		return

	}

}
