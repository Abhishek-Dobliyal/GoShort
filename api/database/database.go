package database

import (
	"context"
	
	"github.com/go-redis/redis/v8"
	"goshort/helpers"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: helpers.DbAddress,
		Password: helpers.DbPass,
		DB: dbNo,
	})

	return rdb
}