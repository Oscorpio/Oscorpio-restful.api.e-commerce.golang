package usecase

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
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

func (m *memberUsecase) CreateUser(ctx context.Context, params *domain.User) error {
	user, _ := m.memberRepo.GetUser(ctx, params.Email)
	if user != nil {
		return domain.ErrConflict
	}

	salt := getSalt()
	np := params.Password + salt

	hashedP, hashErr := hashByBcrypt(np)
	if hashErr != nil {
		return hashErr
	}
	params.Salt = salt
	params.Password = hashedP

	err := m.memberRepo.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (m *memberUsecase) Login(ctx context.Context, email, pwd string) (string, error) {
	user, err := m.memberRepo.GetUser(ctx, email)
	if err != nil {
		return "", err
	}

	np := pwd + user.Salt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(np))
	if err != nil {
		return "", domain.ErrForbidden
	}

	uuid := uuid.NewString()
	return uuid, nil
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

func hashByBcrypt(s string) (string, error) {
	cost, convErr := strconv.Atoi(os.Getenv("HASH_COST"))
	if convErr != nil {
		log.Fatal("env HASH_COST must be integer")
	}

	hashedP, hashErr := bcrypt.GenerateFromPassword([]byte(s), cost)
	if hashErr != nil {
		return "", hashErr
	}

	return string(hashedP), nil
}
