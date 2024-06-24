package db

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Database struct {
	*redis.Client
	logger *zap.Logger
}

func OpenDatabase(l *zap.Logger, client *redis.Client) *Database {
	return &Database{
		client,
		l,
	}
}
