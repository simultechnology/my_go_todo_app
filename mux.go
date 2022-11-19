package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json: charset=utf-8")
		// To avoid static analysis errors, discard this return value intentionally
		_, _ = w.Write([]byte(`{"status": "oK"}`))
	})
	return mux
}
