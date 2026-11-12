package helpers

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MigrateContainer struct {
	testcontainers.Container
}

func CreateMigrationsContainer(ctx context.Context, connStr string, opts ...testcontainers.ContainerCustomizer) (*MigrateContainer, error) {
	defaultOpts := []testcontainers.ContainerCustomizer{
		testcontainers.WithCmd(
			// TODO switch to netAlias
			fmt.Sprintf("-database=%s", connStr),
			"-path=/migrations/",
			"up",
		),
		testcontainers.WithWaitStrategy(wait.ForExit()),
	}

	allOpts := append(defaultOpts, opts...)

	mc, err := testcontainers.Run(ctx,
		"migrate/migrate:v4.20.1",
		allOpts...,
	)

	if err != nil {
		return nil, err
	}

	return &MigrateContainer{
		Container: mc,
	}, nil
}
