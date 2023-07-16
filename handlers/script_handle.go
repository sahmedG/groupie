package ConcertAPI

import (
	"net/http"
)

func ScriptsServ(w http.ResponseWriter, req *http.Request, filename string) {
	switch filename {
	case "rendering_data.js":
		http.ServeFile(w, req, "../html/js/rendering_data.js")
	case "manipulation_css.js":
		http.ServeFile(w, req, "../html/js/css_manipulation.js")
	case "search_bands.js":
		http.ServeFile(w, req, "../html/js/search_bands.js")
	case "locations.js":
		http.ServeFile(w, req, "../html/js/locations.js")
	case "filter_bands.js":
		http.ServeFile(w, req, "../html/js/filter_bands.js")
	}
}
