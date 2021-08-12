package usecase

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"restful.api.e-commerce.golang/domain"
)

type memberUsecase struct {
	memberRepo domain.MemberRepo
}

func NewMemberUsecase(dm domain.MemberRepo) domain.MemberUsecase {
	return &memberUsecase{
		memberRepo: dm,
	}
}

func (m *memberUsecase) CreateUser(ctx context.Context, params *domain.CreateUserParams) error {
	salt := getSalt()
	np := []byte(params.Password + salt)
	cost, convErr := strconv.Atoi(os.Getenv("HASH_COST"))
	if convErr != nil {
		log.Fatal("env HASH_COST must be integer")
	}

	hashedP, hashErr := bcrypt.GenerateFromPassword(np, cost)
	if hashErr != nil {
		return hashErr
	}
	params.Salt = salt
	params.Password = string(hashedP)

	err := m.memberRepo.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func getSalt() string {
	salts := make([]byte, 7)
	t := time.Now().UnixNano()
	rand.Seed(t)

	for k := range salts {
		salts[k] = byte(rand.Intn(93) + 33)
	}

	return string(salts)
}
