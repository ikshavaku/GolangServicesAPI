package postgres

import (
	"github.com/google/wire"
)

var PostgresConnectionSet = wire.NewSet(
	ProvidePostgresConfig,
	ProvidePostgresURL,
	ProvidePostgresConnection,
)

var PostgresProviderSet = wire.NewSet(
	PostgresConnectionSet,
	ProvidePostgresQueries,
)
