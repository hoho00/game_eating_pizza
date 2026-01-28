package handlers

import (
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

// RegisterRequest는 회원가입 요청 구조체입니다
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest는 로그인 요청 구조체입니다
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register는 회원가입을 처리합니다
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
			"details": err.Error(),
		})
		return
	}

	// TODO: AuthService.Register 구현 필요
	player, err := h.authService.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"player": player,
	})
}

// Login은 로그인을 처리합니다
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
			"details": err.Error(),
		})
		return
	}

	// TODO: AuthService.Login 구현 필요
	token, player, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"player": player,
	})
}

// RefreshToken은 토큰을 갱신합니다
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: 토큰 갱신 로직 구현
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}
