package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountDelivery struct {
}

func NewAccountDelivery() *accountDelivery {
	return &accountDelivery{}
}

func (h *accountDelivery) Handler(c *gin.Engine) {
	c.Handle(http.MethodPost, "/login", loginHandler())
}
