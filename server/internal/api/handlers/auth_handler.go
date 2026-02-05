package handlers

import (
	"game_eating_pizza/internal/api/dto"
	"game_eating_pizza/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler는 인증 관련 핸들러입니다
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler는 새로운 AuthHandler를 생성합니다
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register 회원가입
// @Summary      회원가입
// @Description  새로운 플레이어를 등록합니다
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterRequest  true  "회원가입 정보"
// @Success      201      {object}  map[string]interface{}  "회원가입 성공"
// @Failure      400      {object}  map[string]interface{}  "잘못된 요청"
// @Failure      500      {object}  map[string]interface{}  "서버 오류"
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	player, err := h.authService.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to register",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"player":  player,
	})
}

// Login 로그인
// @Summary      로그인
// @Description  플레이어 로그인을 처리하고 JWT 토큰을 반환합니다
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "로그인 정보"
// @Success      200      {object}  map[string]interface{}  "로그인 성공"
// @Failure      400      {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	token, player, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid credentials",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"player": player,
	})
}

// RefreshToken 토큰 갱신
// @Summary      토큰 갱신
// @Description  JWT 토큰을 갱신합니다
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RefreshTokenRequest  false  "토큰 갱신 (바디 없음)"
// @Success      200      {object}  map[string]interface{}  "토큰 갱신 성공"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Router       /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: 토큰 갱신 로직 구현 (요청 파라미터: dto.RefreshTokenRequest)
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}
