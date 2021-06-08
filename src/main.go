package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// directory of static files to be served
	dir := "./views/"

	r := mux.NewRouter()
	// This serves files at http://localhost:3000/static/helloworld.HTML
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	log.Fatal(srv.ListenAndServe())
}
