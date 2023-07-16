package ConcertAPI

import "net/http"

func Serve404Html(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "../html/404/404.html")
}

func Serve400Html(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "../html/404/400.html")
}

func Serve500Html(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "../html/404/500.html")
}