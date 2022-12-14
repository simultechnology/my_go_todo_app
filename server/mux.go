package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/simultechnology/my_go_todo_app/handler"
	"github.com/simultechnology/my_go_todo_app/store"
)

func NewMux() http.Handler {
	// mux := http.NewServeMux()
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json: charset=utf-8")
		// To avoid static analysis errors, discard this return value intentionally
		_, _ = w.Write([]byte(`{"status": "oK"}`))
	})
	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServerHTTP)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux
}
