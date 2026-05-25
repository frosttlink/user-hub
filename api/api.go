package api

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db *sql.DB) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", CreateUser(db))
		r.Get("/", FindAll(db))
		r.Get("/{id}", FindById(db))
		r.Put("/{id}", Update(db))
		r.Delete("/{id}", Delete(db))
	})

	return r
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal json data", "error", err)

		http.Error(
			w,
			`{"error":"internal server error"}`,
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response to client", "error", err)
	}
}
