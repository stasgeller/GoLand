package controller

import (
	"api.com/logger"
	"encoding/json"
	"net/http"

	"api.com/parsers"
)

func ParseXML(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("file")
	if err != nil {
		logger.WriteLog(w, err, http.StatusBadRequest)
		return
	}

	defer f.Close()

	users, err := parsers.NewXmlParser().Parse(f)
	if err != nil {
		logger.WriteLog(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
