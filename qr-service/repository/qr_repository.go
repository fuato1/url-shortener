package repository

import (
	"context"
	"log"
	"os"

	"github.com/fuato1/shorturl/qr-service/model"
	"github.com/go-redis/redis/v8"
)

type QrRepository struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewQrRepository() *QrRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	repo := &QrRepository{
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

func (r *QrRepository) GetAll() ([]model.QR, error) {
	return nil, nil
}

func (r *QrRepository) Add(model.QR) error {
	return nil
}
