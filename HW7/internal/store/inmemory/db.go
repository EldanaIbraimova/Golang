package inmemory

import (
	"context"
	"errors"
	"fmt"
	"HW7/internal/models"
	"HW7/internal/store"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	factsRepo *mongo.Collection
	//profilesRepo *mongo.Collection

	mu *sync.RWMutex
}

func (db *DB) Create(ctx context.Context, fact *models.Fact) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, err := db.factsRepo.InsertOne(ctx, fact)
	return err
}

func (db *DB) All(ctx context.Context, filter *models.FactsFilter) ([]*models.Fact, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	//filter := bson.D{{}}

	return db.filterTasks(ctx, filter)
}

func (db *DB) ByID(ctx context.Context, id string) (*models.Fact, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("conversion of id from string to ObjectID: %s", id)
	}
	filter := bson.D{primitive.E{Key: "_id", Value: idObj}}
	u := &models.Fact{}
	ok := db.factsRepo.FindOne(ctx, filter).Decode(u)
	if ok != nil {
		return nil, fmt.Errorf("no user with id %s", id)
	}
	return u, nil
}

func (db *DB) Update(ctx context.Context, fact *models.Fact) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	filter := bson.D{primitive.E{Key: "_id", Value: fact.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "title", Value: fact.Title},
		primitive.E{Key: "categories", Value: fact.Categories},
		primitive.E{Key: "Text", Value: fact.Text},
	}}}

	u := &models.Fact{}

	return db.factsRepo.FindOneAndUpdate(ctx, filter, update).Decode(u)
}

func (db *DB) Delete(ctx context.Context, id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("conversion of id from string to ObjectID: %s", id)
	}
	filter := bson.D{primitive.E{Key: "_id", Value: idObj}}

	res, err := db.factsRepo.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no tasks were deleted")
	}

	return nil
}

func Init() store.FactsRepository {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return &DB{
		factsRepo: client.Database("FactsDB").Collection("facts"),
		mu: new(sync.RWMutex),
	}
}

func (db *DB) filterTasks(ctx context.Context, filter interface{}) ([]*models.Fact, error) {
	var facts []*models.Fact

	cur, err := db.factsRepo.Find(ctx, filter)
	if err != nil {
		return facts, err
	}

	for cur.Next(ctx) {
		var t models.Fact
		err := cur.Decode(&t)
		if err != nil {
			return facts, err
		}

		facts = append(facts, &t)
	}

	if err := cur.Err(); err != nil {
		return facts, err
	}

	cur.Close(ctx)

	if len(facts) == 0 {
		return facts, mongo.ErrNoDocuments
	}

	return facts, nil
}