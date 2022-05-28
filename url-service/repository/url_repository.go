package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fuato1/shorturl/url-service/model"
	"github.com/go-redis/redis/v8"
)

const cacheDuration = 6 * time.Hour

type UrlRepository struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewUrlRepository() *UrlRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	repo := &UrlRepository{
		ctx: context.Background(),
	}

	msg, err := redisClient.Ping(repo.ctx).Result()
	if err != nil {
		log.Printf("unable to connect to redis server: %v", err.Error())
	}

	log.Printf("redis started {%s}", msg)
	repo.redisClient = redisClient

	return repo
}

func (r *UrlRepository) GetAll() (map[string]string, error) {
	urls := map[string]string{}
	iter := r.redisClient.Scan(r.ctx, 0, "", 0).Iterator()

	for iter.Next(r.ctx) {
		value, err := r.redisClient.Get(r.ctx, iter.Val()).Result()
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

func (r *UrlRepository) Add(url model.ShortUrl) error {
	err := r.redisClient.Set(r.ctx, url.URL, url.Source, cacheDuration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *UrlRepository) Get(shortUrl string) (string, error) {
	result, err := r.redisClient.Get(r.ctx, shortUrl).Result()

	if err != nil {
		return "", err
	}

	return result, nil
}
