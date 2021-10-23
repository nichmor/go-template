package db

import (
	"context"
	"encoding/json"
	"fmt"
	setting "go-template/pkg/settings"

	"github.com/go-redis/redis/v8"
)

var RedisConnection *redis.Client

func ConnectRedis() {
	RedisConnection = redis.NewClient(&redis.Options{
		Addr: setting.RedisSetting.Address,
		DB:   setting.RedisSetting.Database,
	})
}

func DisconnectRedis() {
	err := RedisConnection.Close()
	if err != nil {
		fmt.Println("Error while closign connection to redis", err)
	}
}

func Set(c *redis.Client, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(context.TODO(), key, p, 0).Err()
}

func Get(c *redis.Client, key string, dest *interface{}) error {
	result, err := c.Get(context.TODO(), key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(result), dest)

	if err != nil {
		return err
	}

	return nil
}
