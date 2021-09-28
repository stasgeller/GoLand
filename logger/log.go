package logger

import (
	"fmt"
	"net/http"
)

//TODO: Rewrite to separate error-handler package
func WriteLog(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, err.Error())
}
