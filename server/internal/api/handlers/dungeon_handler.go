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

// GetDungeons 던전 목록 조회
// @Summary      던전 목록 조회
// @Description  활성화된 던전 목록을 조회합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  map[string]interface{}  "던전 목록"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Failure      500      {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons [get]
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

// GetDungeon 던전 상세 조회
// @Summary      던전 상세 조회
// @Description  특정 던전의 상세 정보를 조회합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "던전 ID"
// @Success      200  {object}  models.Dungeon  "던전 정보"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      404  {object}  map[string]interface{}  "던전을 찾을 수 없음"
// @Router       /dungeons/{id} [get]
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

// EnterDungeon 던전 입장
// @Summary      던전 입장
// @Description  던전에 입장합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "던전 ID"
// @Success      200  {object}  map[string]interface{}  "입장 성공"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      500  {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons/{id}/enter [post]
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

// ClearDungeon 던전 클리어
// @Summary      던전 클리어
// @Description  던전을 클리어하고 보상을 획득합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "던전 ID"
// @Success      200  {object}  map[string]interface{}  "클리어 성공"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Router       /dungeons/{id}/clear [post]
func (h *DungeonHandler) ClearDungeon(c *gin.Context) {
	// TODO: 구현 필요
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}
