package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	WriteJSONResponse(w, status, map[string]string{"error": err.Error()})
}

func ParseJson(r *http.Request, decoderType any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request Body.")
	}

	return json.NewDecoder(r.Body).Decode(&decoderType)
}
