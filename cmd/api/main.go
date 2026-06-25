package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"ci-cd-go-learn/internal/database"
	"ci-cd-go-learn/internal/handler"
	"ci-cd-go-learn/internal/repository"
	"ci-cd-go-learn/internal/router"
	"ci-cd-go-learn/internal/service"
)

func main() {
	addr := env("ADDR", ":8080")
	databaseURL := env("DATABASE_URL", "postgres://postgres:password@localhost:5432/ci_cd_learn?sslmode=disable")

	ctx := context.Background()
	pool, err := database.NewPool(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	srv := &http.Server{
		Addr:         addr,
		Handler:      router.New(userHandler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("server listening on %s", addr)
	log.Fatal(srv.ListenAndServe())
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}