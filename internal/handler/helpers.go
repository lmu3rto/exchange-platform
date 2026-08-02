package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return ErrBodyEmpty
		}
		return fmt.Errorf("decode json %w", err)
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, dst any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(dst)

	if err != nil {
		return fmt.Errorf("write json %w", err)
	}
	return nil
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})

	if err != nil {
		log.Printf("failed to encode error response: %v", err)
	}
}
