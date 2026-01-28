package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler는 헬스 체크 핸들러입니다
type HealthHandler struct{}

// NewHealthHandler는 새로운 HealthHandler를 생성합니다
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health는 서버 상태를 확인합니다
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}
