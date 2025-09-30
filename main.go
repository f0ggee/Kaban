package main

import (
	"Kaban/Controller"
	"log/slog"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "Service/Fronted/Maine.html")
			return
		}
		http.ServeFile(w, r, "./fronted"+r.URL.Path)
	})

	http.HandleFunc("/maine/api", Controller.Get_From)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("Err cant' do this", "err", err)
		return
	} else {
		slog.Info("Start")
	}

}
