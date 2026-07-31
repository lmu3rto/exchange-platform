package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func decodeJSON(r http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("body is empty")
		}
		return err
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, dst any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(dst)
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}
