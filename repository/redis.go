package repository

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisRepo struct {
	client *redis.Client
}

var repo *redisRepo
var once sync.Once
var ctx = context.Background()

const cacheDuration = 6 * time.Hour

func Get() Repository {
	once.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: "",
			DB:       0,
		})

		repo = &redisRepo{
			client: redisClient,
		}

		_, err := repo.client.Ping(ctx).Result()
		if err != nil {
			log.Printf("unable to connect to redis server: %v", err.Error())
		}
	})

	return repo
}

func (r *redisRepo) GetAll() (map[string]string, error) {
	urls := map[string]string{}
	iter := r.client.Scan(ctx, 0, "", 0).Iterator()

	for iter.Next(ctx) {
		value, err := r.client.Get(ctx, iter.Val()).Result()
		if err != nil {
			return urls, err
		}

		urls[iter.Val()] = value
	}

	if err := iter.Err(); err != nil {
		return urls, err
	}

	return urls, nil
}

func (r *redisRepo) Add(id, source string) error {
	err := r.client.Set(ctx, id, source, cacheDuration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) Get(id string) (string, error) {
	result, err := r.client.Get(ctx, id).Result()

	if err != nil {
		return "", err
	}

	return result, nil
}
