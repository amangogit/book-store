package main

import (
	"log"
	"net/http"

	"github.com/amangogit/book-store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9010", r))
}
