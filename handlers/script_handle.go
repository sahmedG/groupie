package ConcertAPI

import (
	"fmt"
	"net/http"
)

func ScriptsServ(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/js/core.js")
}

func ScriptsServ2(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/js/beautify.js")

}
func ScriptsServ3(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/js/search.js")
}
func ScriptsServ4(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/js/locations.js")
}

func ScriptsServ5(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/js/filter.js")
}