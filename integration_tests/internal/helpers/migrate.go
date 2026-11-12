package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MigrationContainer struct {
	testcontainers.Container
}

func NewMigrationsContainer(ctx context.Context, connStr string, opts ...testcontainers.ContainerCustomizer) (*MigrationContainer, func(), error) {
	defaultOpts := []testcontainers.ContainerCustomizer{
		testcontainers.WithCmd(
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
		return nil, nil, err
	}

	cleanup := func() {
		termCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := mc.Terminate(termCtx); err != nil {
			log.Printf("migration test container terminating error: %v", err)
		}
	}

	state, err := mc.State(ctx)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	if state.ExitCode != 0 {
		cleanup()
		return nil, nil, fmt.Errorf("migrations failed with status code %d", state.ExitCode)
	}

	return &MigrationContainer{
		Container: mc,
	}, cleanup, nil
}
