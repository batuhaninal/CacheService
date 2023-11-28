package repository

import (
	"CacheService/internal/cache/domain/models"
	"CacheService/internal/cache/domain/ports"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) ports.ICacheRepository {
	return &redisRepository{client: client}
}

func (rr redisRepository) CreateOrUpdate(ctx context.Context, model models.CacheModel) {
	err := rr.client.Set(ctx, model.Key, model.Value, time.Hour*10).Err()
	if err != nil {
		panic(err.Error())
	}
}

func (rr redisRepository) Remove(ctx context.Context, key string) {
	rr.client.Del(ctx, key)
}

func (rr redisRepository) Get(ctx context.Context, key string) models.CacheModel {
	value := rr.client.Get(ctx, key)

	/*if value.Err() != nil {
		panic(value.Err())
	}*/

	data, _ := value.Bytes()

	if len(data) > 0 {
		return models.CacheModel{Key: key, Value: data}
	}

	return models.CacheModel{}
}
