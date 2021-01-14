package utils

import "github.com/go-redis/redis"

const (
	LoginKey = "LoginKey:"
)

var redisdb *redis.Client

func init() {
	redisdb = redis.NewClient(&redis.Options{ //return Client
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedisDb() *redis.Client {
	return redisdb
}

func CheckExist(key string) bool {
	cmd := redisdb.Get(key)
	if cmd.Err() == redis.Nil {
		return false
	}
	if cmd.Val() != "" {
		return true
	} else {
		return false
	}
}
