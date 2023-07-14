package main

import (
	ConcertAPI "ConcertAPI/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"text/template"
	"time"
)

var indexTpl *template.Template
var tpl404 *template.Template

func init() {
	// indexTpl = template.Must(template.ParseGlob("templates/index/*.html"))
	// tpl404 = template.Must(template.ParseGlob("templates/404/*.html"))

	timeToWait := 30

	tStart := time.Now()
	var wg sync.WaitGroup
	waitCh := make(chan struct{})
	wg.Add(1)
	log.Printf("Parsing started, if something goes wrong, program will terminate in %v seconds.", timeToWait)
	go func() {
		go func() {
			ConcertAPI.Parse()
			wg.Done()
		}()
		wg.Wait()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		elapsed := time.Since(tStart)
		log.Printf("Parsing took %.4fs\n", elapsed.Seconds())
		log.Println("Init complete.")
	case <-time.After(time.Duration(timeToWait) * time.Second):
		log.Printf("Parsing failed, terminating\n")
		os.Exit(1)
	}
}

func main() {
	var r router
	
	http.ListenAndServe(":8080", &r)
}

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	switch req.URL.Path {
	case "/":
		ConcertAPI.Index(w, req)
	case "/output.css": // Serves the css file
		ConcertAPI.StylesServ(w, req)
	case "/static/vid.mp4":
		ConcertAPI.VidServe(w, req)
	case "/bands":
		ConcertAPI.Bands(w, req)
	case "/artists":
		ConcertAPI.GetArtists(w, req)
	case "/artisttemplate":
		ConcertAPI.ArtistTemplate(w, req)
	case "/filter":
		ConcertAPI.FilterArtists(w,req)
	case "/geocode":
		ConcertAPI.GetGeocode(w,req)
	case "/find":
		ConcertAPI.FindArtist(w,req)
	case "/static/js/core.js":
		ConcertAPI.ScriptsServ(w, req)
	case "/static/js/locations.js":
		ConcertAPI.ScriptsServ2(w, req)
	case "/static/js/search.js":
		ConcertAPI.ScriptsServ3(w, req)
	case "/static/js/beautify.js":
		ConcertAPI.ScriptsServ4(w, req)
	case "/static/js/filter.js":
		ConcertAPI.ScriptsServ5(w, req)
	case "/static":
	 	ConcertAPI.StaticFiles(w,req)
	case "/static/css/sidenav.css":
		ConcertAPI.StylesServ5(w,req)
	case "/static/css/hero.css":
		ConcertAPI.StylesServ3(w,req)
	case "/static/css/card.css":
		ConcertAPI.StylesServ4(w,req)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}
