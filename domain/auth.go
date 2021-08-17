package domain

import "context"

type RedisAuthRepo interface {
	ValidateToken(ctx context.Context, token string) error
}

type AuthUsecase interface {
	ValidateToken(ctx context.Context, authHeader string) error
}
