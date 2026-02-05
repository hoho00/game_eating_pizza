package params

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAuthenticatedPlayerID는 인증 미들웨어에서 설정한 userID를 조회합니다.
// 인증되지 않은 경우 (0, false)를 반환합니다.
func GetAuthenticatedPlayerID(c *gin.Context) (uint, bool) {
	userIDVal, exists := c.Get("userID")
	if !exists || userIDVal == nil {
		return 0, false
	}
	s, ok := userIDVal.(string)
	if !ok {
		return 0, false
	}
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, false
	}
	return uint(id), true
}
