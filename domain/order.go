package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ProductId primitive.ObjectID `json:"productId" bson:"productId" binding:"required"`
	Quantity  int                `json:"qty" bson:"quantity" binding:"required"`
	Price     int                `json:"price,omitempty" bson:"price,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Color     string             `json:"color" bson:"color" binding:"required"`
	Size      string             `json:"size" bson:"size" binding:"required"`
}

type Order struct {
	Item   []*Item `json:"item" bson:"item" binding:"required"`
	Total  int     `json:"total,omitempty" bson:"total,omitempty"`
	Owner  string  `json:"owner" bson:"owner" binding:"required"`
	Addr   string  `json:"addr" bson:"addr" binding:"required"`
	Status bool    `json:"status" bson:"status"`
}

type MongoOrderRepo interface {
	CreateOrder(ctx context.Context, order *Order) error
	ListOrder(ctx context.Context, owner string) (*Order, error)
}

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *Order) error
	ListOrder(ctx context.Context, owner string) (*Order, error)
}
