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
	"github.com/idbeholdv18/up-code/app/hash"
	"github.com/idbeholdv18/up-code/app/mail"
	"github.com/idbeholdv18/up-code/app/token"
	"github.com/idbeholdv18/up-code/app/user"
	"github.com/jackc/pgx/v5/pgxpool"

	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

func main() {
	ctx := context.Background()

	postgresUsername := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresUsername == "" || postgresPassword == "" || postgresDB == "" {
		log.Fatalf("postgres username, password and db must be provided by env")
	}

	jwtAccessSecretKey := os.Getenv("JWT_ACCESS_SECRET")
	if jwtAccessSecretKey == "" {
		log.Fatalf("jwt access token must be provided by env")
	}
	jwtConfirmationSecretKey := os.Getenv("JWT_CONFIRMATION_SECRET")
	if jwtConfirmationSecretKey == "" {
		log.Fatalf("jwt confirmation token must be provided by env")
	}

	rabbitmqUser := os.Getenv("RABBITMQ_DEFAULT_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_DEFAULT_PASS")
	if rabbitmqUser == "" || rabbitmqPassword == "" {
		log.Panicf("rabbitmq user and password must be provided by env")
	}

	connString := fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", postgresUsername, postgresPassword, postgresDB)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("pgxpool creating error: %v", err)
	}

	brokerURI := fmt.Sprintf("amqp://%s:%s@rabbitmq/", rabbitmqUser, rabbitmqPassword)

	env := rmq.NewEnvironment(brokerURI, nil)
	conn, err := env.NewConnection(ctx)
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %v", err)
	}
	defer func() {
		_ = env.CloseConnections(context.Background())
	}()

	const confirmationQueueName = "mail.confirmation"

	_, err = conn.Management().DeclareQueue(ctx, &rmq.QuorumQueueSpecification{
		Name: confirmationQueueName,
	})
	if err != nil {
		log.Fatalf("falied to declare a queue: %v", err)
	}

	publisher, err := conn.NewPublisher(ctx, &rmq.QueueAddress{
		Queue: confirmationQueueName,
	}, nil)
	if err != nil {
		log.Fatalf("failed to create publisher: %v", err)
	}
	defer func() {
		_ = publisher.Close(context.Background())
	}()

	mailPublisher := mail.NewRmqPublisher(publisher)
	accessTokenService := token.NewJwtAccessTokenService([]byte(jwtAccessSecretKey), time.Minute*5)
	confirmationTokenService := token.NewJwtConfirmationTokenService([]byte(jwtConfirmationSecretKey), time.Hour*24)
	passwordHashService := hash.NewArgon2idService(4)

	userRepo := user.NewPgxpoolRepo(pool)
	userService := user.NewBaseService(userRepo)
	authHandler := auth.NewHttpHandler(
		userService,
		accessTokenService,
		passwordHashService,
		confirmationTokenService,
		mailPublisher,
		"http://localhost:8080",
	)

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
