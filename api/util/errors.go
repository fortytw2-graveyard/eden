package util

import (
	"encoding/json"
	"net/http"
)

// JSONError writes the given error and http statuscode to the writer
func JSONError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
