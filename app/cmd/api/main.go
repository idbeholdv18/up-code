package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/idbeholdv18/up-code/app/auth"
	"github.com/idbeholdv18/up-code/app/token"
	"github.com/idbeholdv18/up-code/app/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	postgresUsername := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresUsername == "" || postgresPassword == "" || postgresDB == "" {
		log.Fatalf("postgres username, password and db must be provided by env")
	}

	jwtSecretKey := os.Getenv("JWT_ACCESS_SECRET")
	if jwtSecretKey == "" {
		log.Fatalf("jwt access token must be provided by env")
	}

	connString := fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", postgresUsername, postgresPassword, postgresDB)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("pgxpool creating error: %v", err)
	}

	accessTokenService := token.NewJwtAccessTokenService([]byte(jwtSecretKey), time.Minute*5)

	userRepo := user.NewPgxpoolRepo(pool)
	userService := user.NewBaseService(userRepo)
	authHandler := auth.NewHttpHandler(userService, accessTokenService)

	jwtMiddleware := auth.JwtMiddleware(accessTokenService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.HandleRegister)
		r.Post("/login", authHandler.HandleLogin)
	})

	r.Route("/protected", func(r chi.Router) {
		r.Use(jwtMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			id, _ := auth.GetUserIDFromCtx(r.Context())
			fmt.Fprintf(w, "protected: %d\n", id)
		})
	})

	http.ListenAndServe(":8080", r)
}
