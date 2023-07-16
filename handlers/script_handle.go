package ConcertAPI

import (
	"net/http"
)

func ScriptsServ(w http.ResponseWriter, req *http.Request, filename string) {
	switch filename {
	case "core.js":
		http.ServeFile(w, req, "../html/js/core.js")
	case "beautify.js":
		http.ServeFile(w, req, "../html/js/beautify.js")
	case "search.js":
		http.ServeFile(w, req, "../html/js/search.js")
	case "locations.js":
		http.ServeFile(w, req, "../html/js/locations.js")
	case "filter.js":
		http.ServeFile(w, req, "../html/js/filter.js")
	}
}
