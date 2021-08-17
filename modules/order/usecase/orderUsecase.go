package usecase

import (
	"context"

	"restful.api.e-commerce.golang/domain"
)

type orderUsecase struct {
	mongoOrderRepo   domain.MongoOrderRepo
	mongoProductRepo domain.MongoProductRepo
}

func NewOrderUsecase(
	dmo domain.MongoOrderRepo,
	dmp domain.MongoProductRepo,
) domain.OrderUsecase {
	return &orderUsecase{
		mongoOrderRepo:   dmo,
		mongoProductRepo: dmp,
	}
}

func (o *orderUsecase) CreateOrder(ctx context.Context, order *domain.Order) error {
	var total int
	l := len(order.Item)

	for i := 0; i < l; i++ {
		item := order.Item[i]

		item, price, err := getProduct(o, ctx, item)
		if err != nil {
			return err
		}
		total += price

		up, err := getProductDetail(o, ctx, item)
		if err != nil {
			return err
		}
		if item.Quantity > up.Stock {
			return domain.ErrParamInput
		}

		up.Stock -= item.Quantity
		err = updateStock(o, ctx, up)
		if err != nil {
			return err
		}
	}

	order.Total = total
	err := o.mongoOrderRepo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderUsecase) ListOrder(ctx context.Context, order string) (
	*domain.Order, error,
) {
	r, err := o.mongoOrderRepo.ListOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func getProduct(o *orderUsecase, ctx context.Context, item *domain.Item) (
	*domain.Item, int, error,
) {
	p, err := o.mongoProductRepo.ListProductById(ctx, item.ProductId)
	if err != nil {
		return nil, 0, err
	}

	item.Name = p.Name
	item.Price = p.Price

	return item, p.Price * item.Quantity, nil
}

func getProductDetail(o *orderUsecase, ctx context.Context, item *domain.Item) (
	*domain.Detail, error,
) {
	up, err := o.mongoProductRepo.ListDetail(
		ctx,
		&domain.Detail{
			ProductId: &item.ProductId,
			Size:      item.Size,
			Color:     item.Color,
		},
	)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func updateStock(o *orderUsecase, ctx context.Context, up *domain.Detail) error {
	err := o.mongoProductRepo.UpdateDetail(ctx, up)
	if err != nil {
		return err
	}

	return nil
}
