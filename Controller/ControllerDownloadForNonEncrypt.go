package Controller

import (
	"Kaban/Service/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

func CookiGetInControllerDownloaderNoEnc(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		http.Error(w, "cookie dont sen", http.StatusUnauthorized)
		return false
	}

	_, ok := session.Values["cookie"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)

	}

	return true
}
func getNameFromUrl2(r *http.Request) string {
	vars := mux.Vars(r)

	name := vars["name"]
	slog.Info(name)
	return name

}
func DownloadWithNonEcnrypt(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Status method don't allow", http.StatusBadRequest)
		return
	}

	name := getNameFromUrl2(r)

	ok := CookiGetInControllerDownloaderNoEnc(w, r)
	if !ok {
		return
	}

	err := Handlers.SDownloadWithNonEncrypt(w, name)
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
