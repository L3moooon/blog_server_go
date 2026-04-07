package core

import (
	"blog_backend_go/global"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.Addr, redisCfg.Port),
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("redis连接错误: %v", err))
	}
	log.Println("✅ redis连接成功")
	return client
}
