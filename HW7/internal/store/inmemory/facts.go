package inmemory

import (
	"context"
	"fmt"
	"HW7/internal/models"
	"sync"
)

type FactsRepo struct {
	data map[int]*models.Fact

	mu *sync.RWMutex
}

func (db *FactsRepo) Create(ctx context.Context, fact *models.Fact) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[fact.ID] = fact
	return nil
}

func (db *FactsRepo) All(ctx context.Context) ([]*models.Fact, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	facts := make([]*models.Fact, 0, len(db.data))
	for _, fact := range db.data {
		facts = append(facts, fact)
	}

	return facts, nil
}

func (db *FactsRepo) ByID(ctx context.Context, id int) (*models.Fact, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	fact, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No fact with id %d", id)
	}

	return fact, nil
}

func (db *FactsRepo) Update(ctx context.Context, fact *models.Fact) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[fact.ID] = fact
	return nil
}

func (db *FactsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
