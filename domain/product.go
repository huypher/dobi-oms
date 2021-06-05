package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int    `json:"id"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}

type ProductUsecase interface {
	Add(ctx context.Context, sku string, name string, description string, categoryID int) error
}

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
}

type ProductDelivery interface {
	Handler(c *gin.RouterGroup)
}
