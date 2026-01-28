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

// GetMe 내 정보 조회
// @Summary      내 정보 조회
// @Description  현재 로그인한 플레이어의 정보를 조회합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  models.Player  "플레이어 정보"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Failure      404      {object}  map[string]interface{}  "플레이어를 찾을 수 없음"
// @Router       /players/me [get]
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

// UpdateMe 내 정보 수정
// @Summary      내 정보 수정
// @Description  현재 로그인한 플레이어의 정보를 수정합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  map[string]interface{}  "수정 성공"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Router       /players/me [put]
func (h *PlayerHandler) UpdateMe(c *gin.Context) {
	// TODO: 구현 필요
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}

// GetLeaderboard 리더보드 조회
// @Summary      리더보드 조회
// @Description  레벨이 높은 상위 플레이어 목록을 조회합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        limit   query     int  false  "조회할 플레이어 수 (기본값: 10)"
// @Success      200     {object}  map[string]interface{}  "리더보드"
// @Failure      500     {object}  map[string]interface{}  "서버 오류"
// @Router       /players/leaderboard [get]
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
