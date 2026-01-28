package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler는 에러를 처리하는 미들웨어입니다
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 에러가 있는지 확인
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Printf("Error: %v", err.Error())

			// 에러 타입에 따라 적절한 응답
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
				"message": err.Error(),
			})
		}
	}
}
