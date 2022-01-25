package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fibonacci", fibHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "FIBONACCI ENDPOINT")
}