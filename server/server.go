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
	http.ListenAndServe(":8000", &r)
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
	case "/bands.html":
		ConcertAPI.Bands(w, req)
	case "/artists":
		ConcertAPI.GetArtists(w, req)
	case "/artisttemplate.html":
		ConcertAPI.ArtistTemplate(w, req)
	case "/core.js":
		ConcertAPI.ScriptsServ(w, req)

	default:
		http.Error(w, "404 Not Found", 404)
	}
}
