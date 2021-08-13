package domain

import "context"

type User struct {
	UserName string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required,min=6,max=10"`
	Salt     string `bson:"salt"`
	Addr     string `json:"addr" bson:"addr" binding:"required"`
	Email    string `json:"email" bons:"email" binding:"required,email"`
}

type MongoRepo interface {
	CreateUser(ctx context.Context, params *User) error
	GetUser(ctx context.Context, email string) (*User, error)
}

type RedisRepo interface {
	StoreToken(ctx context.Context, token, email string) error
	DeleteToken(ctx context.Context, token string) error
}

type MemberUsecase interface {
	CreateUser(ctx context.Context, params *User) error
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, token string) error
}
