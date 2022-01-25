package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HELLO THIS IS WORKING")
}