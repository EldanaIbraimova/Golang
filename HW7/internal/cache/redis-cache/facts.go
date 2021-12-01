package redis_cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"HW7/internal/cache"
	"HW7/internal/models"
	"time"
)

func (rc RedisCache) Facts() cache.FactsCacheRepo {
	if rc.facts == nil {
		rc.facts = newFactsRepo(rc.client, rc.expires)
	}

	return rc.facts
}

type FactsRepo struct {
	client  *redis.Client
	expires time.Duration
}

func newFactsRepo(client *redis.Client, exp time.Duration) cache.FactsCacheRepo {
	return &FactsRepo{
		client:  client,
		expires: exp,
	}
}

func (c FactsRepo) Set(ctx context.Context, key string, value []*models.Fact) error {
	factsBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = c.client.Set(ctx, key, factsBytes, c.expires*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (c FactsRepo) Get(ctx context.Context, key string) ([]*models.Fact, error) {
	result, err := c.client.Get(ctx, key).Result()
	switch err {
	case nil:
		break
	case redis.Nil:
		return nil, nil
	default:
		return nil, err
	}

	facts := make([]*models.Fact, 0)
	if err = json.Unmarshal([]byte(result), &facts); err != nil {
		return nil, err
	}

	return facts, nil
}
