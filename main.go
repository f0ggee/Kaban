package main

import (
	"Kaban/Controller"
	"Kaban/Service/Handlers"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "Service/Fronted/Maine.html")

		}
	})

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Service/Fronted/login.html")
	})
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Service/Fronted/Register.html")
	})
	router.HandleFunc("/main", func(writer http.ResponseWriter, request *http.Request) {

		http.ServeFile(writer, request, "Service/Fronted/Main_Page.html")

	})
	router.HandleFunc("/wait", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "Service/Fronted/WaitDownload.html")

	})

	router.HandleFunc("/URL/{name}", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "Service/Fronted/UrlFronted.html")

	}).Name("fileName")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла", err)

	}

	router.HandleFunc("/login/api", Controller.Loging).Methods("POST")
	router.HandleFunc("/register/api", Controller.Controller_Register).Methods("POST")
	router.HandleFunc("/d/{name}", func(writer http.ResponseWriter, request *http.Request) {
		var p Handlers.CustomError

		ch := make(chan string)
		Dow := p.ServiceDownload(ch, writer, request)
		if Dow.Err != nil {
			slog.Info(Dow.Message)
			return
		}

		Handlers.Delete(ch)

	}).Methods(http.MethodGet)

	router.HandleFunc("/downloader/api", func(writer http.ResponseWriter, request *http.Request) {

		Controller.ControlerFileUploader(writer, request, router)

	}).Methods(http.MethodPost)
	router.HandleFunc("/maine/api", func(writer http.ResponseWriter, request *http.Request) {
		Controller.Get_From(writer, request)

	}).Methods("GET")
	router.HandleFunc("/doUrl/api", func(writer http.ResponseWriter, request *http.Request) {

		NameFile := Controller.CUrlUp(writer, request)
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(map[string]string{
			"Url": "http://localhost:8080/" + "d/" + NameFile,
		}); err != nil {
			slog.Error("Json can't be treated -", err)
			return
		}
	}).Methods(http.MethodGet)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		slog.Error("Err cant' do this", "err", err)
		return
	} else {
		slog.Info("Start")
	}

}
