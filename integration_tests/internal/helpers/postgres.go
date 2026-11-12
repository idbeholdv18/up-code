package helpers

import (
	"context"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func (c *PostgresContainer) InternalConnectionString(alias string) (string, error) {
	u, err := url.Parse(c.ConnectionString)
	if err != nil {
		return "", err
	}

	u.Host = net.JoinHostPort(alias, "5432")
	return u.String(), nil
}

func NewPostgresContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*PostgresContainer, func(), error) {
	defaultOpts := []testcontainers.ContainerCustomizer{
		// postgres.WithInitScripts(filepath.Join("..", "..", "test_data", "init_db.sql")),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(30 * time.Second),
		),
	}

	allOpts := append(defaultOpts, opts...)

	pgc, err := postgres.Run(ctx,
		"postgres:18",
		allOpts...,
	)

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		termCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := pgc.Terminate(termCtx); err != nil {
			log.Printf("postgres test container terminating error: %v", err)
		}
	}

	connStr, err := pgc.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgc,
		ConnectionString:  connStr,
	}, cleanup, nil
}
