package config

import (
	"os"

	"github.com/go-redis/redis"
)

// Redis Global Redis cliend
var Redis = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_SERVER"),
	Password: os.Getenv("REDIS_PASSWORD"),
})
