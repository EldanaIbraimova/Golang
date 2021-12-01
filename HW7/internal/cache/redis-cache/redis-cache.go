package redis_cache
import (
	"HW7/internal/cache"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	client *redis.Client
	expires time.Duration
	facts   cache.FactsCacheRepo
}


func NewRedisCache(host string, db int, exp time.Duration) cache.Cache {
	c := new(RedisCache)
	c.client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       db,
	})
	c.expires = exp

	return c
}

func (rc *RedisCache) Close() error {
	if err := rc.client.Close(); err != nil {
		return err
	}

	return nil
}

func (rc *RedisCache) DeleteAll(ctx context.Context) error {
	if err := rc.client.FlushAll(ctx).Err(); err != nil {
		return err
	}

	return nil
}