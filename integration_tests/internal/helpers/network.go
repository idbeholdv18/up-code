package helpers

import (
	"context"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
)

func NewNetwork(ctx context.Context) (*testcontainers.DockerNetwork, func(), error) {
	net, err := network.New(ctx)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		rmCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := net.Remove(rmCtx); err != nil {
			log.Printf("network remove error: %v", err)
		}
	}

	return net, cleanup, nil
}
