package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pghuy/dobi-oms/auth"
)

func loginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.PostForm("username")
		password := c.PostForm("password")
		token, err := auth.Login(c, userName, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ERROR",
				"error": map[string]string{
					"code":    "INVALID_REQUEST",
					"message": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"token":  token,
		})
	}
}
