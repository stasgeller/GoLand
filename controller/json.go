package controller

import (
	"api.com/logger"
	"api.com/parsers"
	"encoding/json"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("file")
	if err != nil {
		logger.WriteLog(w, err, http.StatusBadRequest)
		return
	}

	defer f.Close()

	users, err := parsers.NewJsonParser().Parse(f)
	if err != nil {
		logger.WriteLog(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
