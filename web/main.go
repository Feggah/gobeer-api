package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/feggah/gobeer-api/core/beer"
	"github.com/feggah/gobeer-api/web/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data/beer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := beer.NewService(db)
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeBeerHandlers(r, n, service)
	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
