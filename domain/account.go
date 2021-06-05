package domain

import "github.com/gin-gonic/gin"

type Account struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AccountRepository interface {
}

type AccountDelivery interface {
	Handler(c *gin.Engine)
}
