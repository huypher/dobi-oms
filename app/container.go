package app

import (
	"net/http"

	"github.com/pghuy/dobi-oms/product/delivery"

	"github.com/google/wire"
	accountDelivery "github.com/pghuy/dobi-oms/auth/delivery"
	"github.com/pghuy/dobi-oms/domain"
	"github.com/pghuy/dobi-oms/infra"
	productRepository "github.com/pghuy/dobi-oms/product/repository"
	productUsecase "github.com/pghuy/dobi-oms/product/usecase"
	"github.com/uptrace/bun"
)

var set = wire.NewSet(
	ProvideConfig,
	ProvidePostgres,
	ProvideProductDelivery,
	ProvideProductUsecase,
	ProvideProductRepository,
	ProvideHttpHandler,
	ProvideRestService,
	ProvideAccountDelivery,
)

func ProvideConfig() (*infra.Config, error) {
	return infra.NewConfig()
}

func ProvidePostgres(cfg *infra.Config) (*bun.DB, func(), error) {
	return infra.NewPostgres(cfg)
}

func ProvideProductDelivery(
	productUsecase domain.ProductUsecase,
) domain.ProductDelivery {
	return delivery.NewProductDelivery(productUsecase)
}

func ProvideProductUsecase(productRepository domain.ProductRepository) (domain.ProductUsecase, error) {
	return productUsecase.NewProductUsecase(productRepository)
}

func ProvideProductRepository(db *bun.DB) (domain.ProductRepository, error) {
	return productRepository.NewProductRepository(db)
}

func ProvideHttpHandler(productDelivery domain.ProductDelivery, accountDelivery domain.AccountDelivery) http.Handler {
	return infra.NewHttpHandler(productDelivery, accountDelivery)
}

func ProvideRestService(httpHandler http.Handler) *http.Server {
	return infra.NewRestService(httpHandler)
}

func ProvideAccountDelivery() domain.AccountDelivery {
	return accountDelivery.NewAccountDelivery()
}
