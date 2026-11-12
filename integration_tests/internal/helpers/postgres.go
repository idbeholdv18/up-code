package helpers

import (
	"context"
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

	u.Host = alias
	return u.String(), nil
}

func CreatePostgresContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*PostgresContainer, error) {
	defaultOpts := []testcontainers.ContainerCustomizer{
		// postgres.WithInitScripts(filepath.Join("..", "..", "test_data", "init_db.sql")),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5 * time.Second),
		),
	}

	allOpts := append(defaultOpts, opts...)

	pgc, err := postgres.Run(ctx,
		"postgres:18",
		allOpts...,
	)

	if err != nil {
		return nil, err
	}

	// TODO fix conn str to container net alias
	connStr, err := pgc.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgc,
		ConnectionString:  connStr,
	}, nil
}
