package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

type Db struct {
	Pool  *pgxpool.Pool
	Mongo *mongo.Client
}

// databaseConfig is the configuration for the database package.
type databaseConfig struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func newConfigMongoDB() *databaseConfig {
	return &databaseConfig{
		host:       os.Getenv("MONGODB_HOST"),
		database:   os.Getenv("MONGODB_DATABASE"),
		password:   os.Getenv("MONGODB_PASSWORD"),
		user:       os.Getenv("MONGODB_USER"),
		ctxTimeout: 60 * time.Second,
	}
}

func newConfigPostgres() *databaseConfig {
	return &databaseConfig{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
