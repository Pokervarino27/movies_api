package util

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

//Conn interface with all methods of redis that we will use
type Conn interface {
	SetKey(key string, data interface{}, ttl time.Duration) bool
	GetKey(key string, data interface{}) (bool, error)
	Delete(key string) (bool, error)
}

var (
	ctx = context.Background()
)

type redisClient struct {
	client *redis.Client
}

//RedisInit initiliaze the connection with redis
func RedisInit() (Conn, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "ec2-54-193-123-106.us-west-1.compute.amazonaws.com:6379",
		Password: "486a247ed44697be66127749a86e7efb3a11f46c",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return &redisClient{client: client}, nil
}

func (c *redisClient) GetKey(key string, data interface{}) (bool, error) {

	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		err = json.Unmarshal([]byte(val), &data)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (c *redisClient) SetKey(key string, data interface{}, ttl time.Duration) bool {

	err := c.client.Set(ctx, key, data, ttl*time.Second).Err()
	if err != nil {
		return false
	}
	return true
}

func (c *redisClient) Delete(key string) (bool, error) {

	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
