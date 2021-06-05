package infra

import (
	"net/http"

	"github.com/pghuy/dobi-oms/auth"

	"github.com/pghuy/dobi-oms/domain"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler(
	productDelivery domain.ProductDelivery,
	accountHandler domain.AccountDelivery,
) http.Handler {
	router := gin.Default()

	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(healthcheck.Default())
	router.Use(gin.Recovery())

	accountHandler.Handler(router)

	v1 := router.Group("/v1")
	{
		v1.Use(auth.Middleware())
		productDelivery.Handler(v1)
	}

	return router
}
