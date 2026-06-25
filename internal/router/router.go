package router

import (
	"net/http"

	"ci-cd-go-learn/internal/handler"
)

func New(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("GET /api/users", userHandler.ListUsers)
	mux.HandleFunc("POST /api/users", userHandler.CreateUser)

	return mux
}