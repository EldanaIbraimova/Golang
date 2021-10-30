package inmemory

import (
"HW7/internal/models"
"HW7/internal/store"
"sync"
)

type DB struct {
	factsRepo    store.FactsRepository
	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}


func (db *DB) Facts() store.FactsRepository {
	if db.factsRepo == nil {
		db.factsRepo = &FactsRepo{
			data: make(map[int]*models.Fact),
			mu:   new(sync.RWMutex),
		}
	}

	return db.factsRepo
}
