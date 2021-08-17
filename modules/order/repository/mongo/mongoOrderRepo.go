package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"restful.api.e-commerce.golang/domain"
)

type mongoOrderRepo struct {
	db *mongo.Database
}

func NewMongoOrderRepo(db *mongo.Database) domain.MongoOrderRepo {
	return &mongoOrderRepo{
		db,
	}
}

func (m *mongoOrderRepo) CreateOrder(ctx context.Context, order *domain.Order) error {
	coll := m.db.Collection("order")

	_, err := coll.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoOrderRepo) ListOrder(ctx context.Context, owner string) (*domain.Order, error) {
	coll := m.db.Collection("order")
	r := &domain.Order{}
	filter := bson.M{
		"owner": owner,
	}

	err := coll.FindOne(ctx, filter).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
