package controller

import (
	"api.com/parsers"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {
	f, h, err := r.FormFile("file")

	if err != nil {
		panic(fmt.Sprintf("ERROR [form-file]: %s", err))
	}

	users, err := parsers.NewJsonParser().Parse(f, *h)

	if err != nil {
		panic(fmt.Sprintf("ERROR [parse-json]: %s", err))
	}

	json.NewEncoder(w).Encode(users)
}