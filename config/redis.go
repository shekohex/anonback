package config

import "github.com/go-redis/redis"

// Redis Global Redis cliend
var Redis = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
