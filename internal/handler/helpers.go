package handler

import (
	"encoding/json"
	"net/http"
)

func decodeJSON(r http.Request, dst any) (error) {
	return json.NewDecoder(r.Body).Decode(dst)
}

func writeJSON(w http.ResponseWriter, status int, dst any) (error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(dst)
}

func writeError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}