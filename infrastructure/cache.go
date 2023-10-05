package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/riyan-eng/golang-boilerplate-one/env"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnRedis() {
	addr := fmt.Sprintf("%v:%v", env.REDIS_HOST, env.REDIS_PORT)
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: env.REDIS_USERNAME,
		Password: env.REDIS_PASSWORD,
		DB:       env.REDIS_DATABASE,
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		fmt.Printf("redis: can't ping to redis - %v \n", err)
		os.Exit(1)
	}
	fmt.Println("redis: connection opened")
}
