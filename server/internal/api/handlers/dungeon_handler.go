package handlers

import (
	"game_eating_pizza/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DungeonHandler는 던전 관련 핸들러입니다
type DungeonHandler struct {
	dungeonService *services.DungeonService
}

// NewDungeonHandler는 새로운 DungeonHandler를 생성합니다
func NewDungeonHandler(dungeonService *services.DungeonService) *DungeonHandler {
	return &DungeonHandler{
		dungeonService: dungeonService,
	}
}

// GetDungeons는 던전 목록을 조회합니다
func (h *DungeonHandler) GetDungeons(c *gin.Context) {
	dungeons, err := h.dungeonService.GetActiveDungeons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get dungeons",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dungeons": dungeons,
	})
}

// GetDungeon는 특정 던전 정보를 조회합니다
func (h *DungeonHandler) GetDungeon(c *gin.Context) {
	dungeonIDStr := c.Param("id")
	dungeonID, err := strconv.ParseUint(dungeonIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dungeon ID",
		})
		return
	}

	dungeon, err := h.dungeonService.GetDungeonByID(uint(dungeonID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Dungeon not found",
		})
		return
	}

	c.JSON(http.StatusOK, dungeon)
}

// EnterDungeon는 던전에 입장합니다
func (h *DungeonHandler) EnterDungeon(c *gin.Context) {
	dungeonIDStr := c.Param("id")
	dungeonID, err := strconv.ParseUint(dungeonIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dungeon ID",
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	playerID, err := strconv.ParseUint(userID.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	err = h.dungeonService.EnterDungeon(uint(playerID), uint(dungeonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to enter dungeon",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Entered dungeon successfully",
	})
}

// ClearDungeon는 던전을 클리어합니다
func (h *DungeonHandler) ClearDungeon(c *gin.Context) {
	// TODO: 구현 필요
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}
