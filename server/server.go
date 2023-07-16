package main

import (
	ConcertAPI "ConcertAPI/handlers"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	ConcertAPI.Parse()
	var r router
	http.ListenAndServe(":8080", &r)
}

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "/" {
		ConcertAPI.Index(w, req)
	} else if path == "/bands" {
		ConcertAPI.Bands(w, req)
	} else if path == "/templates" {
		ConcertAPI.ArtistTemplate(w, req)
	} else if path == "/filter" {
		ConcertAPI.FilterBands(w, req)
	} else if path == "/geocode" {
		ConcertAPI.FetchLocCode(w, req)
	} else if path == "/artists" {
		ConcertAPI.GetBands(w, req)
	} else if path == "/find" {
		ConcertAPI.FindBand(w, req)
	} else if strings.Contains(path, "/static") {
		filename := strings.ReplaceAll(req.URL.Path, "/static/", "")
		fmt.Println(filename)
		ConcertAPI.StaticFiles(w, req, filename)
	} else if strings.Contains(path, "/css") {
		filename := strings.ReplaceAll(req.URL.Path, "/css/", "")
		fmt.Println(filename)
		ConcertAPI.StylesServ(w, req, filename)
	} else if strings.Contains(path, "/js") {
		filename := strings.ReplaceAll(req.URL.Path, "/js/", "")
		fmt.Println(filename)
		ConcertAPI.ScriptsServ(w, req, filename)
	} else {
		ConcertAPI.Serve404Html(w, req)
	}
}
