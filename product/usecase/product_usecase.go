package usecase

import (
	"context"

	"github.com/pghuy/dobi-oms/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
}

func NewProductUsecase(productRepository domain.ProductRepository) (*productUsecase, error) {
	return &productUsecase{
		productRepository: productRepository,
	}, nil
}

func (u *productUsecase) Add(
	ctx context.Context,
	sku string,
	name string,
	description string,
	categoryID int,
) error {
	return u.productRepository.Save(ctx, &domain.Product{
		Sku:         sku,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	})
}
