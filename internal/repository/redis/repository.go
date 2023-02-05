package redis

import (
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{
		db: db,
	}
}
