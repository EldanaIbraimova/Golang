package store

import (
	"context"
	"HW7/internal/models"
)

type Store interface {
	Facts() FactsRepository
}

type FactsRepository interface {
	Create(ctx context.Context, fact *models.Fact) error
	All(ctx context.Context) ([]*models.Fact, error)
	ByID(ctx context.Context, id int) (*models.Fact, error)
	Update(ctx context.Context, game *models.Fact) error
	Delete(ctx context.Context, id int) error
}
