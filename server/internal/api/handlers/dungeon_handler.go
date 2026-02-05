package handlers

import (
	"game_eating_pizza/internal/api/dto"
	"game_eating_pizza/internal/models"
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

// GetAllDungeons 전체 던전 목록 조회
// @Summary      전체 던전 목록 조회
// @Description  전체 던전 목록을 조회합니다 (활성/비활성 포함)
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  map[string]interface{}  "던전 목록"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Failure      500      {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons/all [get]
func (h *DungeonHandler) GetAllDungeons(c *gin.Context) {
	dungeons, err := h.dungeonService.GetAllDungeons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get dungeons",
		})
		return
	}

	// DTO로 변환
	responses := make([]dto.DungeonResponse, len(dungeons))
	for i, dungeon := range dungeons {
		responses[i] = dto.DungeonResponse{
			ID:         dungeon.ID,
			Name:       dungeon.Name,
			Type:       dungeon.Type,
			Difficulty: dungeon.Difficulty,
			IsActive:   dungeon.IsActive,
			StartTime:  dungeon.StartTime,
			EndTime:    dungeon.EndTime,
			CreatedAt:  dungeon.CreatedAt,
			UpdatedAt:  dungeon.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"dungeons": responses,
	})
}

// GetActiveDungeons 활성 던전 목록 조회
// @Summary      활성 던전 목록 조회
// @Description  활성화된 던전 목록만 조회합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  map[string][]dto.DungeonResponse  "던전 목록"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Failure      500      {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons/active [get]
func (h *DungeonHandler) GetActiveDungeons(c *gin.Context) {
	dungeons, err := h.dungeonService.GetActiveDungeons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get active dungeons",
		})
		return
	}

	// DTO로 변환
	responses := make([]dto.DungeonResponse, len(dungeons))
	for i, dungeon := range dungeons {
		responses[i] = dto.DungeonResponse{
			ID:         dungeon.ID,
			Name:       dungeon.Name,
			Type:       dungeon.Type,
			Difficulty: dungeon.Difficulty,
			IsActive:   dungeon.IsActive,
			StartTime:  dungeon.StartTime,
			EndTime:    dungeon.EndTime,
			CreatedAt:  dungeon.CreatedAt,
			UpdatedAt:  dungeon.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"dungeons": responses,
	})
}

// CreateDungeon 던전 생성
// @Summary      던전 생성
// @Description  새로운 던전을 생성합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.CreateDungeonRequest  true  "던전 생성 정보"
// @Success      201   {object}  dto.DungeonResponse  "생성된 던전 정보"
// @Failure      400   {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401   {object}  map[string]interface{}  "인증 실패"
// @Failure      500   {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons [post]
func (h *DungeonHandler) CreateDungeon(c *gin.Context) {
	var req dto.CreateDungeonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	dungeon := &models.Dungeon{
		Name:       req.Name,
		Type:       req.Type,
		Difficulty: req.Difficulty,
		IsActive:   isActive,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	}
	if err := h.dungeonService.CreateDungeon(dungeon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create dungeon"})
		return
	}

	response := dto.DungeonResponse{
		ID:         dungeon.ID,
		Name:       dungeon.Name,
		Type:       dungeon.Type,
		Difficulty: dungeon.Difficulty,
		IsActive:   dungeon.IsActive,
		StartTime:  dungeon.StartTime,
		EndTime:    dungeon.EndTime,
		CreatedAt:  dungeon.CreatedAt,
		UpdatedAt:  dungeon.UpdatedAt,
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateDungeon 던전 수정
// @Summary      던전 수정
// @Description  던전 정보를 수정합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int  true  "던전 ID"
// @Param        body  body      dto.UpdateDungeonRequest  true  "수정할 던전 정보"
// @Success      200   {object}  dto.DungeonResponse  "수정된 던전 정보"
// @Failure      400   {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401   {object}  map[string]interface{}  "인증 실패"
// @Failure      404   {object}  map[string]interface{}  "던전을 찾을 수 없음"
// @Failure      500   {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons/{id} [put]
func (h *DungeonHandler) UpdateDungeon(c *gin.Context) {
	dungeonIDStr := c.Param("id")
	dungeonID, err := strconv.ParseUint(dungeonIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dungeon ID"})
		return
	}

	dungeon, err := h.dungeonService.GetDungeonByID(uint(dungeonID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dungeon not found"})
		return
	}

	var req dto.UpdateDungeonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	if req.Name != nil {
		dungeon.Name = *req.Name
	}
	if req.Type != nil {
		dungeon.Type = *req.Type
	}
	if req.Difficulty != nil {
		dungeon.Difficulty = *req.Difficulty
	}
	if req.IsActive != nil {
		dungeon.IsActive = *req.IsActive
	}
	if req.StartTime != nil {
		dungeon.StartTime = req.StartTime
	}
	if req.EndTime != nil {
		dungeon.EndTime = req.EndTime
	}

	if err := h.dungeonService.UpdateDungeon(dungeon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update dungeon"})
		return
	}

	response := dto.DungeonResponse{
		ID:         dungeon.ID,
		Name:       dungeon.Name,
		Type:       dungeon.Type,
		Difficulty: dungeon.Difficulty,
		IsActive:   dungeon.IsActive,
		StartTime:  dungeon.StartTime,
		EndTime:    dungeon.EndTime,
		CreatedAt:  dungeon.CreatedAt,
		UpdatedAt:  dungeon.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteDungeon 던전 삭제
// @Summary      던전 삭제
// @Description  던전을 삭제합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "던전 ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      404  {object}  map[string]interface{}  "던전을 찾을 수 없음"
// @Failure      500  {object}  map[string]interface{}  "서버 오류"
// @Router       /dungeons/{id} [delete]
func (h *DungeonHandler) DeleteDungeon(c *gin.Context) {
	dungeonIDStr := c.Param("id")
	dungeonID, err := strconv.ParseUint(dungeonIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dungeon ID"})
		return
	}

	if _, err := h.dungeonService.GetDungeonByID(uint(dungeonID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dungeon not found"})
		return
	}

	if err := h.dungeonService.DeleteDungeon(uint(dungeonID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete dungeon"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetDungeon 던전 상세 조회
// @Summary      던전 상세 조회
// @Description  특정 던전의 상세 정보를 조회합니다
// @Tags         dungeons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "던전 ID"
// @Success      200  {object}  dto.DungeonResponse  "던전 정보"
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

	// DTO로 변환
	response := dto.DungeonResponse{
		ID:         dungeon.ID,
		Name:       dungeon.Name,
		Type:       dungeon.Type,
		Difficulty: dungeon.Difficulty,
		IsActive:   dungeon.IsActive,
		StartTime:  dungeon.StartTime,
		EndTime:    dungeon.EndTime,
		CreatedAt:  dungeon.CreatedAt,
		UpdatedAt:  dungeon.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
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
