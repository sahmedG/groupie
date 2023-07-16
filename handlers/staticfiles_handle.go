package ConcertAPI

import (
	"net/http"
)

func StaticFiles(w http.ResponseWriter, req *http.Request,filename string) {
	switch filename{
	case "vid.mp4":
		http.ServeFile(w, req, "../html/static/vid.mp4")
	case "sad.png":
		http.ServeFile(w, req, "../html/static/sad.png")
	case "floyd.jpeg":
		http.ServeFile(w, req, "../html/static/floyd.css")
	case "concert.jpeg":
		http.ServeFile(w, req, "../html/static/concert.jpeg")
	case "calendar.svg":
		http.ServeFile(w, req, "../html/static/calendar.svg")
	case "favicon.ico":
		http.ServeFile(w, req, "../html/static/favicon.ico")
	}

}