package ConcertAPI

import (
	"net/http"
)

 func StylesServ(w http.ResponseWriter, req *http.Request,filename string) {
	switch filename{
	case "output.css":
		http.ServeFile(w, req, "../html/css/output.css")
	case "hero.css":
		http.ServeFile(w, req, "../html/css/hero.css")
	case "card.css":
		http.ServeFile(w, req, "../html/css/card.css")
	case "sidenav.css":
		http.ServeFile(w, req, "../html/css/sidenav.css")
	case "main.css":
		http.ServeFile(w, req, "../html/css/main.css")
	case "404.css":
		http.ServeFile(w, req, "../html/css/404.css")
	}
 }
