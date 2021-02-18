package utils

import (
	"rickyWeb/src/utils/viperUtils"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func init() {
	//redis initial , setting can see redis.go opt.init()
	redisdb = redis.NewClient(&redis.Options{ //return Client
		Addr:     viperUtils.GetCommonParams("redisPath"),
		Password: "",
		DB:       0,
	})

	cmd := redisdb.Ping()
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
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

func GetDel(key string) (jsonStr string) {
	cmd := redisdb.Get(key)
	if cmd.Val() != "" {
		jsonStr = cmd.Val()
		redisdb.Del(key)
	}
	return
}
