package ConcertAPI

import (
	"fmt"
	"html/template"
	"net/http"
)

var indexTpl *template.Template

// func Bands(w http.ResponseWriter, req *http.Request) {
// 	origin := req.Header.Get("Origin")
// 	w.Header().Set("Access-Control-Allow-Origin", origin)
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	http.ServeFile(w, req, "../html/bands.html")
// }

func Bands(w http.ResponseWriter, r *http.Request) {
	indexTpl = template.Must(template.ParseGlob("../templates/index/*.html"))
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "static/assets/favicon.ico")
	} else if r.URL.Path == "/bands" {
		switch r.Method {
		case "GET":
			fmt.Println("here")
			indexTpl.ExecuteTemplate(w, "bands.html", nil)
		default:
			CallErrorPage(w, r, 405)
		}
	} else {
		CallErrorPage(w,r,404)
	} 
}

func CallErrorPage(w http.ResponseWriter, r *http.Request, errorCode int) {
	fmt.Println("url"+r.URL.Path)
	switch errorCode {
	case 404:
		r.URL.Path = "/"
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w,r,"../templates/404/404.html")
	case 405:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
