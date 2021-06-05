package http_response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrInternal     = "có lỗi hệ thống xảy ra"
	ErrUnauthorized = "chưa đăng nhập"
)

func Error(c *gin.Context, err error) {
	var httpCode int
	var message string
	switch err.Error() {
	case ErrUnauthorized:
		httpCode = http.StatusUnauthorized
		message = ErrUnauthorized
	default:
		httpCode = http.StatusInternalServerError
		message = ErrInternal
	}

	c.JSON(httpCode, gin.H{
		"message": message,
	})
}

func Abort(c *gin.Context, err error) {
	var httpCode int
	var message string
	switch err.Error() {
	case ErrUnauthorized:
		httpCode = http.StatusUnauthorized
		message = ErrUnauthorized
	default:
		httpCode = http.StatusInternalServerError
		message = ErrInternal
	}

	c.AbortWithStatusJSON(httpCode, gin.H{
		"message": message,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "thành công",
	})
}
