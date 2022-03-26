package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fuato1/shorturl/internal/domain/url"

	"github.com/go-redis/redis/v8"
)

const CacheDuration = 6 * time.Hour

type Repo struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewRepo() *Repo {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	repo := &Repo{
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

func (r *Repo) GetAll() ([]url.ShortUrl, error) {
	var urls []url.ShortUrl
	iter := r.redisClient.Scan(r.ctx, 0, "", 0).Iterator()

	for iter.Next(r.ctx) {
		result, err := r.redisClient.Get(r.ctx, iter.Val()).Result()
		if err != nil {
			return urls, err
		}

		url := url.ShortUrl{
			Source: result,
			URL:    iter.Val(),
		}
		urls = append(urls, url)
	}

	if err := iter.Err(); err != nil {
		return urls, err
	}

	return urls, nil
}

func (r *Repo) Add(url url.ShortUrl) error {
	err := r.redisClient.Set(r.ctx, url.URL, url.Source, CacheDuration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUrl(shortUrl string) (string, error) {
	result, err := r.redisClient.Get(r.ctx, shortUrl).Result()

	if err != nil {
		return "", err
	}

	return result, nil
}
