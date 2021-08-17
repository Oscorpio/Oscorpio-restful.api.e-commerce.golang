package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"restful.api.e-commerce.golang/domain"
)

type redisAuthRepo struct {
	db *redis.Client
}

func NewRedisAuthRepo(db *redis.Client) domain.RedisAuthRepo {
	return &redisAuthRepo{
		db,
	}
}

func (r *redisAuthRepo) ValidateToken(ctx context.Context, token string) error {
	_, err := r.db.Get(ctx, token).Result()
	if err != nil {
		return err
	}

	return nil
}
