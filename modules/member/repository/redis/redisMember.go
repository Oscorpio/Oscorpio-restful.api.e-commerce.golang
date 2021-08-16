package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"restful.api.e-commerce.golang/domain"
)

type redisMemberRepo struct {
	db *redis.Client
}

func NewRedisMemberRepo(db *redis.Client) domain.RedisMemberRepo {
	return &redisMemberRepo{
		db,
	}
}

func (r *redisMemberRepo) StoreToken(ctx context.Context, token, email string) error {
	err := r.db.SetEX(ctx, token, email, 60*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisMemberRepo) DeleteToken(ctx context.Context, token string) error {
	err := r.db.Del(ctx, token).Err()
	if err != nil {
		return err
	}

	return nil
}
