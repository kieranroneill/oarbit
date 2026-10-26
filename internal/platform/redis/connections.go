package redis

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewConnection(
	ctx context.Context,
	config Config,
) (*redis.Client, error) {
	address := net.JoinHostPort(config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		MinIdleConns: 2,
		PoolSize:     10,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		_ = client.Close()

		return nil, fmt.Errorf("failed to ping redis at %s: %w", address, err)
	}

	return client, nil
}
