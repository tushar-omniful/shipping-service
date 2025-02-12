//go:build wireinject
// +build wireinject

package order_controller

import (
	"context"

	"github.com/google/wire"
	"github.com/omniful/go_commons/db/sql/postgres"
)

func Wire(ctx context.Context, db *postgres.DbCluster, nameSpace string) (*Controller, error) {
	panic(wire.Build(ProviderSet))
}
