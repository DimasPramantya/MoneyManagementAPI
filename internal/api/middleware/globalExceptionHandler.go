package middleware

import (
	"github.com/gin-gonic/gin"
	. "github.com/dimas-pramantya/money-management/internal/domain"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *CustomError:
				c.JSON(e.Code, gin.H{
					"message": e.Message,
					"errors":  e.Errors,
					"code":    e.Code,
				})
			default:
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			c.Abort()
		}
	}
}