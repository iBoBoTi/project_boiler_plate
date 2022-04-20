package database

import (
	"errors"
)

var (
	errInvalidDatabaseInstance = errors.New("invalid db instance")
)

const (
	InstancePostgres int = iota
	InstanceMongo
)

// NewDatabaseFactory returns Db type based of the db instance provided
func NewDatabaseFactory(instance int) (*Db, error) {
	switch instance {
	case InstancePostgres:
		return ConnectPostgres(newConfigPostgres())
	case InstanceMongo:
		return ConnectMongoDB(newConfigMongoDB())
	default:
		return nil, errInvalidDatabaseInstance
	}
}
