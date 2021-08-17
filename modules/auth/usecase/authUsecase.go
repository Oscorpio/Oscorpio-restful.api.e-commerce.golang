package usecase

import (
	"context"
	"strings"

	"restful.api.e-commerce.golang/domain"
)

type authUsecase struct {
	redisAuthRepo domain.RedisAuthRepo
}

func NewAuthUsecase(dr domain.RedisAuthRepo) domain.AuthUsecase {
	return &authUsecase{
		redisAuthRepo: dr,
	}
}

func (a *authUsecase) ValidateToken(ctx context.Context, authHeader string) error {
	if authHeader == "" || len(strings.Split(authHeader, "Bearer ")) < 2 {
		return domain.ErrForbidden
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	err := a.redisAuthRepo.ValidateToken(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
