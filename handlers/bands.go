package ConcertAPI

import (
	"fmt"
	"html/template"
	"net/http"
)

var indexTpl *template.Template
var tpl404 *template.Template

// func Bands(w http.ResponseWriter, req *http.Request) {
// 	origin := req.Header.Get("Origin")
// 	w.Header().Set("Access-Control-Allow-Origin", origin)
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	http.ServeFile(w, req, "../html/bands.html")
// }

func Bands(w http.ResponseWriter, r *http.Request) {
	indexTpl = template.Must(template.ParseGlob("../templates/index/*.html"))
	tpl404 = template.Must(template.ParseGlob("../templates/404/*.html"))
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "static/assets/favicon.ico")
	}
	switch r.Method {
	case "GET":
		fmt.Println("here")
		indexTpl.ExecuteTemplate(w, "bands.html", nil)
	default:
		callErrorPage(w, r, 405)
	}
}

func callErrorPage(w http.ResponseWriter, r *http.Request, errorCode int) {
	var errorMsg string

	switch errorCode {
	case 404:
		w.WriteHeader(http.StatusNotFound)
		errorMsg = "404 Page not found"
	case 405:
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorMsg = "405 Wrong method"
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		errorMsg = "400 Bad request"
	default:
		w.WriteHeader(http.StatusInternalServerError)
		errorMsg = "500 Internal error"
		errorCode = 500
	}

	tpl404.ExecuteTemplate(w, "404.html", struct {
		ErrorCode int
		Error     string
	}{
		ErrorCode: errorCode,
		Error:     errorMsg,
	})
}
