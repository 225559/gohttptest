package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Fatal(http.ListenAndServe(":80", handler()))
}

func handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {}
