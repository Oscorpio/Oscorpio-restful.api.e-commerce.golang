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
		p, err := o.mongoProductRepo.ListProductById(ctx, order.Item[i].ProductId)
		if err != nil {
			return err
		}
		order.Item[i].Name = p.Name
		order.Item[i].Price = p.Price

		up, err := o.mongoProductRepo.ListUnitProduct(
			ctx,
			&domain.UnitProduct{
				ProductId: &order.Item[i].ProductId,
				Size:      order.Item[i].Size,
				Color:     order.Item[i].Color,
			},
		)
		if err != nil {
			return err
		}

		total += p.Price * order.Item[i].Quantity

		if order.Item[i].Quantity > up.Stock {
			return domain.ErrParamInput
		}
		up.Stock -= order.Item[i].Quantity
		err = o.mongoProductRepo.UpdateUnitProduct(ctx, up)
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
