package main

import (
	"log/slog"
	"main/api"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	connStr := "postgres://user:password@localhost:5432/users_db?sslmode=disable"
	db, err := api.NewDB(connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
