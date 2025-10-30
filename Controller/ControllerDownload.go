package Controller

import (
	"Kaban/Service/Handlers"
	"net/http"
)

func Dow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	Handlers.ServiceDownload(w, r)

}
