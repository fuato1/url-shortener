package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheService struct {
	redisClient *redis.Client
}

var (
	cacheService = &CacheService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitCache() *CacheService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	msg, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started: ping msg = {%s}", msg)
	cacheService.redisClient = redisClient

	return cacheService
}

func SaveUrl(shortUrl string, originalUrl string, userId string) {
	err := cacheService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		panic(fmt.Sprintf("Error saving key URL. Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func GetInitialUrl(shortUrl string) string {
	result, err := cacheService.redisClient.Get(ctx, shortUrl).Result()

	if err != nil {
		panic(fmt.Sprintf("Error retrieving URL. Error: %v - shortUrl: %s\n", err, shortUrl))
	}

	return result
}
