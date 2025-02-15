//go:build wireinject
// +build wireinject

package order_controller

import (
	"context"

	"github.com/google/wire"
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/redis"
)

func Wire(ctx context.Context, db *postgres.DbCluster, redisClient *redis.Client) (*Controller, error) {
	panic(wire.Build(ProviderSet))
}
