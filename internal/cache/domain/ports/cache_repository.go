package ports

import (
	"CacheService/internal/cache/domain/models"
	"context"
)

type ICacheRepository interface {
	CreateOrUpdate(ctx context.Context, model models.CacheModel)
	Remove(ctx context.Context, key string)
	Get(ctx context.Context, key string) models.CacheModel
}
