package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/dobi-oms/domain"
)

type productDelivery struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductDelivery(productUsecase domain.ProductUsecase) *productDelivery {
	return &productDelivery{
		ProductUsecase: productUsecase,
	}
}

func (h *productDelivery) Handler(c *gin.RouterGroup) {
	productGroup := c.Group("/product")
	productGroup.Handle(http.MethodPost, "/add", addProduct(h.ProductUsecase))
}
