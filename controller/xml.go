package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api.com/parsers"
)

func ParseXML(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("file")

	if err != nil {
		panic(fmt.Sprintf("ERROR: %s", err))
	}

	users, err := parsers.NewXmlParser().Parse(f)

	if err != nil {
		fmt.Fprintln(w, "ERROR: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(users)
}
