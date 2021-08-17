package domain

import (
	"context"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id     *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string              `json:"name" bson:"name" binding:"required"`
	Image  *Image              `json:"image,omitempty" bson:"image,omitempty"`
	Price  int                 `json:"price" bson:"price" binding:"required"`
	Detail []*Detail           `json:"detail,omitempty" bson:"detail,omitempty"`
}

type Image struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" binding:"required"`
	ProductId primitive.ObjectID `json:"productId" bson:"productId" binding:"required"`
	Name      string             `json:"image,omitempty" bson:"image,omitempty" binding:"required"`
	URL       string             `json:"url,omitempty" bson:"url,omitempty" binding:"required"`
}

type Detail struct {
	ProductId *primitive.ObjectID `json:"productId" bson:"productId" binding:"required"`
	Color     string              `json:"color" bson:"color" binding:"required"`
	Size      string              `json:"size" bson:"size" binding:"required"`
	Stock     int                 `json:"stock" bson:"stock" binding:"required"`
}

type MongoProductRepo interface {
	StoreProduct(ctx context.Context, params *Product) error
	StoreImageInfo(ctx context.Context, image *Image) error
	StoreDetail(ctx context.Context, Detail *Detail) error
	ListProducts(ctx context.Context) ([]*Product, error)
	ListProductById(ctx context.Context, id primitive.ObjectID) (*Product, error)
	ListDetail(ctx context.Context, params *Detail) (*Detail, error)
	UpdateDetail(ctx context.Context, params *Detail) error
}

type ProductUsecase interface {
	CreateProduct(ctx context.Context, params *Product) error
	StoreImage(ctx context.Context, image *multipart.FileHeader, id primitive.ObjectID) error
	StoreDetail(ctx context.Context, detail *Detail) error
	ListProducts(ctx context.Context) ([]*Product, error)
	ListProductById(ctx context.Context, id primitive.ObjectID) (*Product, error)
}
