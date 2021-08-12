package mongo

import (
	"restful.api.e-commerce.golang/domain"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoMemberRepo struct {
	db *mongo.Database
}

func NewMongoMemberRepo(db *mongo.Database) domain.MemberRepo {
	return &mongoMemberRepo{
		db,
	}
}

func (m *mongoMemberRepo) CreateUser(ctx context.Context, params *domain.CreateUserParams) error {
	coll := m.db.Collection("member")

	_, err := coll.InsertOne(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
