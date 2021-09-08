package config

import (
	"log"
	"net/http"

	"api.com/parsers"
	simplerest "api.com/simple-rest"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter().StrictSlash(true)
	r.Headers("Content-Type", "application/json")

	r.HandleFunc("/articles", simplerest.List).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}", simplerest.Fetch).Methods("GET")
	r.HandleFunc("/articles", simplerest.Create).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}", simplerest.Delete).Methods("DELETE")

	r.HandleFunc("/parsers/xml", parsers.ParseXml).Methods("GET")

	log.Fatal(http.ListenAndServe(":1000", r))
}
