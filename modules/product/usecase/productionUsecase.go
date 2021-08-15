package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"restful.api.e-commerce.golang/domain"

	_ "github.com/joho/godotenv/autoload"
)

type productUsecase struct {
	mongoProductRepo domain.MongoProductRepo
}

func NewProductUsecase(dm domain.MongoProductRepo) domain.ProductUsecase {
	return &productUsecase{
		mongoProductRepo: dm,
	}
}

func (p *productUsecase) StoreImage(
	ctx context.Context,
	image *multipart.FileHeader,
	oid primitive.ObjectID,
) (
	string, error) {
	if _, err := os.Stat(os.Getenv("IMAGE_PATH") + image.Filename); err == nil {
		return "", domain.ErrConflict
	}

	dst := os.Getenv("IMAGE_PATH") + image.Filename
	imageUrl := fmt.Sprintf("%s:%s/i/%s",
		os.Getenv("FILE_SERVER_DOMAIN"),
		os.Getenv("FILE_SERVER_PORT"),
		image.Filename,
	)
	f, err := image.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, f)
	if err != nil {
		return "", err
	}

	r, err := p.mongoProductRepo.StoreImageInfo(ctx,
		&domain.Image{Name: image.Filename, URL: imageUrl, ProductId: oid})
	if err != nil {
		return "", err
	}

	return r, nil
}

func (p *productUsecase) CreateProduct(ctx context.Context, dp *domain.Product) error {
	err := p.mongoProductRepo.StoreProduct(ctx, dp)

	if err != nil {
		return err
	}

	return nil
}

func (p *productUsecase) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	r, err := p.mongoProductRepo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (p *productUsecase) StoreUnitStock(ctx context.Context, u *domain.UnitProduct) error {
	err := p.mongoProductRepo.StoreUnitStock(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (p *productUsecase) ListProductById(ctx context.Context, id primitive.ObjectID) (
	*domain.Product, error,
) {
	r, err := p.mongoProductRepo.ListProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	return r, nil
}
