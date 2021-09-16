package app

import (
	"log"
	"net/http"

	"api.com/controller"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter().StrictSlash(true)
	r.Headers("Content-Type", "application/json")

	r.HandleFunc("/articles", controller.List).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}", controller.Fetch).Methods("GET")
	r.HandleFunc("/articles", controller.Create).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}", controller.Delete).Methods("DELETE")

	r.HandleFunc("/parsers/xml", controller.Parse).Methods("GET")

	log.Fatal(http.ListenAndServe(":1000", r))
}