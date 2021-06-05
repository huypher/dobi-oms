package delivery

import (
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/dobi-oms/domain"
	"github.com/pghuy/dobi-oms/pkg/http_response"
)

func addProduct(productUsecase domain.ProductUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				logrus.WithError(err).Error("Add product")
			}
		}()

		var p addProductParam
		err = c.ShouldBindJSON(&p)
		if err != nil {
			http_response.Error(c, err)
			return
		}

		err = productUsecase.Add(c, p.Sku, p.Name, p.Description, p.CategoryID)
		if err != nil {
			http_response.Error(c, err)
			return
		}

		http_response.Success(c)
	}
}
