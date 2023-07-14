package ConcertAPI

import (
	"fmt"
	"net/http"
)

func StylesServ(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/output.css")
}
func StylesServ3(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/css/card.css")
}
func StylesServ4(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/css/hero.css")
}
func StylesServ5(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/css/sidenav.css")
}

func StaticFiles(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../html/static/assets/calendar.svg")
	http.ServeFile(w, req, "../html/static/assets/break.png")
	http.ServeFile(w, req, "../html/static/assets/favicon.ico")
	http.ServeFile(w, req, "../html/static/assets/sad-freddie.png")
}