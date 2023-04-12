package cache

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func (r *Redis) GetClient(address, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return rdb
}
