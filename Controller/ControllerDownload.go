package Controller

import (
	"net/http"
)

func Dow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	//Handlers.pServiceDownload(w, r)

}
