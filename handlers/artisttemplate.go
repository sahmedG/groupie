package ConcertAPI

import (
	"net/http"
)

func ArtistTemplate(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../html/templates/artisttemplate.html")
	w.Header().Set("Content-Type", "text/html")
}
