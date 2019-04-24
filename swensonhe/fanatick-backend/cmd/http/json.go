package main

import (
	"encoding/json"
	"net/http"
)

// JSONWriterFunc is a function that writes JSON.
type JSONWriterFunc func(v interface{}, status int)

// NewJSONWriter returns a new JSON writer.
func NewJSONWriter(w http.ResponseWriter) JSONWriterFunc {
	return func(v interface{}, status int) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(v)
	}
}

// Write writes an interface as JSON.
func (f JSONWriterFunc) Write(v interface{}, status int) {
	f(v, status)
}

// decodeJSON decodes *http.Request body into obj which needs to be a pointer
func decodeJSON(r *http.Request, obj interface{}) error {

	if obj == nil {
		return nil
	}

	return json.NewDecoder(r.Body).Decode(obj)
}
