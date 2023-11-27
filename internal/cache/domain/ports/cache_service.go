package ports

import (
	"CacheService/internal/cache/domain/models"
	"context"
)

type ICacheService interface {
	Save(ctx context.Context, model models.CacheModel)
	Remove(ctx context.Context, key string)
	Get(ctx context.Context, key string) models.CacheModel
}
