package services

import (
	"CacheService/internal/cache/domain/models"
	"CacheService/internal/cache/domain/ports"
	"context"
)

type redisService struct {
	redisRepo ports.ICacheRepository
}

func NewRedisService(redisRepo ports.ICacheRepository) ports.ICacheService {
	return &redisService{redisRepo: redisRepo}
}

func (rs redisService) Save(ctx context.Context, model models.CacheModel) {
	rs.redisRepo.CreateOrUpdate(ctx, model)
}

func (rs redisService) Get(ctx context.Context, key string) models.CacheModel {
	return rs.redisRepo.Get(ctx, key)
}

func (rs redisService) Remove(ctx context.Context, key string) {
	rs.redisRepo.Remove(ctx, key)
}
