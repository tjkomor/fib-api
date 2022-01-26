package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hide"
	dbname   = "fibonacci"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	
	if err != nil {
		panic(err)
	}

	fmt.Println("I AM CONNECTED TO THE FIB DATABASE")

	r := mux.NewRouter()
	r.HandleFunc("/fibonacci", FibHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "FIBONACCI ENDPOINT")
}