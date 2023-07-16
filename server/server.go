package main

import (
	ConcertAPI "ConcertAPI/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

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
	path := req.URL.Path


	if path == "/" {
		ConcertAPI.Index(w, req)
	} else if path == "/bands"{
		ConcertAPI.Bands(w, req)
	} else if path == "/filter" {
		ConcertAPI.FilterArtists(w, req)
	} else if path == "/geocode" {
		ConcertAPI.GetGeocode(w, req)
	} else if path == "/artists"{
		ConcertAPI.GetArtists(w, req)
	} else if path == "/find" {
		ConcertAPI.FindArtist(w,req)	
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
	} 
}
