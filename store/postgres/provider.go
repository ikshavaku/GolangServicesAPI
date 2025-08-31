package postgres

import (
	"context"

	"github.com/ikshavaku/catalogue/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresURL string

func ProvidePostgresConnection(url PostgresURL) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), string(url))
	if err != nil {
		// TODO Think something better to do in network error case
		// This would essential cause the container to crash if connection to postgres isn't established
		panic(err)
	}
	return pool
}

func ProvidePostgresURL(config utils.PostgresConfig) PostgresURL {
	return PostgresURL(config.PostgresConnectionURL())
}

/*
Entry point dependency for postgres
*/
func ProvidePostgresConfig() utils.PostgresConfig {
	return utils.GetConfig().Postgres
}

func ProvidePostgresQueries(connection *pgxpool.Pool) *Queries {
	return New(connection)
}
