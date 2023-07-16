package ConcertAPI

import (

	"html/template"
	"net/http"
)

var indexTpl *template.Template

func Bands(w http.ResponseWriter, r *http.Request) {
	indexTpl = template.Must(template.ParseGlob("../html/templates/*.html"))
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "static/assets/favicon.ico")
	} else if r.URL.Path == "/bands" {
		switch r.Method {
		case "GET":
			indexTpl.ExecuteTemplate(w, "bands.html", nil)
		default:
			Serve400Html(w, r)
		}
	} else {
		Serve404Html(w,r)
	} 
}

