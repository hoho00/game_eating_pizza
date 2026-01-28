package handlers

import (
	"game_eating_pizza/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PlayerHandler는 플레이어 관련 핸들러입니다
type PlayerHandler struct {
	playerService *services.PlayerService
}

// NewPlayerHandler는 새로운 PlayerHandler를 생성합니다
func NewPlayerHandler(playerService *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}

// GetMe는 현재 로그인한 플레이어 정보를 조회합니다
func (h *PlayerHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	// TODO: userID를 uint로 변환하는 로직 필요
	playerID, err := strconv.ParseUint(userID.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	player, err := h.playerService.GetPlayerByID(uint(playerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Player not found",
		})
		return
	}

	c.JSON(http.StatusOK, player)
}

// UpdateMe는 현재 로그인한 플레이어 정보를 업데이트합니다
func (h *PlayerHandler) UpdateMe(c *gin.Context) {
	// TODO: 구현 필요
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}

// GetLeaderboard는 리더보드를 조회합니다
func (h *PlayerHandler) GetLeaderboard(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	players, err := h.playerService.GetTopPlayersByLevel(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get leaderboard",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"players": players,
	})
}
