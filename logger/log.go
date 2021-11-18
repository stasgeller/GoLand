package logger

import (
	"encoding/json"
	"net/http"
)

//TODO: Rewrite to separate error-handler package
func WriteLog(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
