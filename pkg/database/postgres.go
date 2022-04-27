package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	pool *pgxpool.Pool
}

// ConnectPostgres connects to the postgres database pool and assigns it the Db struct pool field returning Db
func ConnectPostgres(c *databaseConfig) (*Db, error) {
	logrus.Info("Connecting to PostgreSQL DB pool")
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
	)
	config, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Db{Pool: pool}, nil
}
