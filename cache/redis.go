package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client *redis.Client
	cache  *cache.Cache
	ctx    context.Context
	ttl    time.Duration
}

func New() Cache {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("unable to connect to redis: %v", err)
	}

	cache := cache.New(&cache.Options{
		Redis:      client,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &redisCache{
		client: client,
		cache:  cache,
		ctx:    ctx,
		ttl:    1 * time.Hour,
	}
}

func (rc *redisCache) All() (map[string]string, error) {
	records := make(map[string]string)
	iter := rc.client.Scan(rc.ctx, 0, "", 0).Iterator()

	for iter.Next(rc.ctx) {
		value, err := rc.client.Get(rc.ctx, iter.Val()).Result()
		if err != nil {
			return records, err
		}

		records[iter.Val()] = value
	}

	if err := iter.Err(); err != nil {
		return records, err
	}

	return records, nil
}

func (rc *redisCache) Add(key, value string) error {
	err := rc.cache.Once(&cache.Item{
		Key: key,
		Do: func(i *cache.Item) (any, error) {
			return value, nil
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (rc *redisCache) Get(key string) (string, error) {
	var record string

	err := rc.cache.Once(&cache.Item{
		Key:   key,
		Value: record,
	})
	if err != nil {
		return record, err
	}

	return record, nil
}
