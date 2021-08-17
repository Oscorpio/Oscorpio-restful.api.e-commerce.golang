package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"restful.api.e-commerce.golang/domain"
)

type mongoProductRepo struct {
	db *mongo.Database
}

func NewMongoProductRepo(db *mongo.Database) domain.MongoProductRepo {
	return &mongoProductRepo{
		db,
	}
}

func (m *mongoProductRepo) StoreImageInfo(ctx context.Context, image *domain.Image) error {
	coll := m.db.Collection("image")
	_, err := coll.InsertOne(ctx, image)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoProductRepo) StoreProduct(ctx context.Context, dp *domain.Product) error {
	coll := m.db.Collection("product")
	_, err := coll.InsertOne(ctx, dp)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoProductRepo) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	r := []*domain.Product{}
	lookupImageStage := bson.D{
		{"$lookup", bson.D{
			{"from", "image"},
			{"localField", "_id"},
			{"foreignField", "productId"},
			{"as", "image"},
		}},
	}
	unwindStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$image"},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"name", 1},
			{"image", 1},
			{"price", 1},
		}},
	}
	coll := m.db.Collection("product")

	curs, err := coll.Aggregate(ctx,
		mongo.Pipeline{
			lookupImageStage,
			unwindStage,
			projectStage,
		})
	if err != nil {
		return nil, err
	}
	defer curs.Close(ctx)

	for curs.Next(ctx) {
		p := &domain.Product{}
		err = curs.Decode(p)
		if err != nil {
			return nil, err
		}

		r = append(r, p)
	}

	return r, nil
}

func (m *mongoProductRepo) StoreDetail(ctx context.Context, u *domain.Detail) error {
	coll := m.db.Collection("detail")
	_, err := coll.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoProductRepo) ListProductById(ctx context.Context, id primitive.ObjectID) (
	*domain.Product, error) {
	r := &domain.Product{}
	matchStage := bson.D{
		{"$match", bson.D{
			{"_id", id},
		}},
	}
	lookupImageStage := bson.D{
		{"$lookup", bson.D{
			{"from", "image"},
			{"localField", "_id"},
			{"foreignField", "productId"},
			{"as", "image"},
		}},
	}
	lookupDetailStage := bson.D{
		{"$lookup", bson.D{
			{"from", "detail"},
			{"localField", "_id"},
			{"foreignField", "productId"},
			{"as", "detail"},
		}},
	}
	unwindStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$image"},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"name", 1},
			{"image", 1},
			{"price", 1},
			{"detail", 1},
		}},
	}
	coll := m.db.Collection("product")

	curs, err := coll.Aggregate(ctx,
		mongo.Pipeline{
			matchStage,
			lookupImageStage,
			lookupDetailStage,
			unwindStage,
			projectStage,
		})
	if err != nil {
		return nil, err
	}
	defer curs.Close(ctx)

	for curs.Next(ctx) {
		err := curs.Decode(r)
		if err != nil {
			return nil, err
		}
	}

	return r, nil

}

func (m *mongoProductRepo) ListDetail(ctx context.Context, params *domain.Detail) (
	*domain.Detail, error,
) {
	r := &domain.Detail{}
	coll := m.db.Collection("detail")
	filter := bson.M{
		"productId": params.ProductId,
		"size":      params.Size,
		"color":     params.Color,
	}

	err := coll.FindOne(ctx, filter).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *mongoProductRepo) UpdateDetail(
	ctx context.Context,
	params *domain.Detail) error {
	coll := m.db.Collection("detail")
	filter := bson.M{
		"productId": params.ProductId,
		"size":      params.Size,
		"color":     params.Color,
	}
	update := bson.M{
		"$set": bson.M{
			"stock": params.Stock,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
