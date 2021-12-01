package store

import (
	"context"
	"HW7/internal/models"
)


type FactsRepository interface {
	Create(ctx context.Context, fact *models.Fact) error
	All(ctx context.Context, filter *models.FactsFilter) ([]*models.Fact, error)
	ByID(ctx context.Context, id string) (*models.Fact, error)
	Update(ctx context.Context, fact *models.Fact) error
	Delete(ctx context.Context, id string) error
}