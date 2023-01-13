package server

import (
	"context"
	"github.com/simultechnology/my_go_todo_app/clock"
	"github.com/simultechnology/my_go_todo_app/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/simultechnology/my_go_todo_app/handler"
	"github.com/simultechnology/my_go_todo_app/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	// mux := http.NewServeMux()
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json: charset=utf-8")
		// To avoid static analysis errors, discard this return value intentionally
		_, _ = w.Write([]byte(`{"status": "oK"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := &store.Repository{
		Clocker: clock.RealClocker{},
	}
	at := &handler.AddTask{DB: db, Repo: r, Validator: v}
	mux.Post("/tasks", at.ServerHTTP)
	lt := &handler.ListTask{DB: db, Repo: r}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux, cleanup, nil
}
