package middleware

import (
	"game_eating_pizza/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware는 JWT 토큰을 검증하는 미들웨어입니다
// 개발 환경이고 SKIP_AUTH=true일 때는 Authorization 헤더 없이도 요청을 허용합니다
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		
		// SKIP_AUTH이 활성화되어 있으면 (개발 환경에서는 기본값)
		if cfg.SkipAuth {
			// 임시 사용자 ID 설정 (테스트용)
			if authHeader != "" {
				// Authorization 헤더가 있으면 토큰에서 추출
				parts := strings.Split(authHeader, " ")
				if len(parts) == 2 && parts[0] == "Bearer" {
					c.Set("userID", parts[1])
				} else {
					c.Set("userID", "dev-user")
				}
			} else {
				// Authorization 헤더가 없으면 기본값 사용
				c.Set("userID", "dev-user")
			}
			c.Next()
			return
		}
		
		// SKIP_AUTH가 비활성화되어 있으면 (프로덕션 환경) 정상적으로 검증
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// "Bearer <token>" 형식에서 토큰 추출
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		token := parts[1]

		// TODO: JWT 토큰 검증 및 사용자 ID 추출
		// 임시로 토큰을 사용자 ID로 사용 (실제로는 JWT 파싱 필요)
		// 예: userID, err := jwt.ValidateToken(token, cfg.JWTSecret)
		
		// 임시 구현: 토큰을 사용자 ID로 사용 (개발용)
		// 프로덕션에서는 반드시 JWT 검증 로직 구현 필요
		c.Set("userID", token) // 임시

		c.Next()
	}
}
