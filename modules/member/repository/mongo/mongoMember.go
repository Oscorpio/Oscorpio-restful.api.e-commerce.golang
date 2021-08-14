package mongo

import (
	"restful.api.e-commerce.golang/domain"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoMemberRepo struct {
	db *mongo.Database
}

func NewMongoMemberRepo(db *mongo.Database) domain.MongoRepo {
	return &mongoMemberRepo{
		db,
	}
}

func (m *mongoMemberRepo) CreateUser(ctx context.Context, params *domain.User) error {
	coll := m.db.Collection("member")

	_, err := coll.InsertOne(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoMemberRepo) GetUser(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	coll := m.db.Collection("member")
	filter := bson.M{"email": email}
	err := coll.FindOne(ctx, filter).Decode(user)

	if err != nil {
		return nil, domain.ErrNotFound
	}

	return user, nil
}
