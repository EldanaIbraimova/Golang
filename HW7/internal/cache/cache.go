package cache

import (
	"context"
	"HW7/internal/models"
)

type Cache interface {
	Close() error

	Facts() FactsCacheRepo

	DeleteAll(ctx context.Context) error
}

type FactsCacheRepo interface {
	Set(ctx context.Context, key string, value []*models.Fact) error
	Get(ctx context.Context, key string) ([]*models.Fact, error)
}