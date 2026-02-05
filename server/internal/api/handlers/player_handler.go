package handlers

import (
	"game_eating_pizza/internal/api/params"
	"game_eating_pizza/internal/api/dto"
	"game_eating_pizza/internal/services"
	"net/http"

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
// @Success      200      {object}  dto.PlayerResponse  "플레이어 정보"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Failure      404      {object}  map[string]interface{}  "플레이어를 찾을 수 없음"
// @Router       /players/me [get]
func (h *PlayerHandler) GetMe(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	req := dto.GetMeRequest{PlayerID: playerID}

	player, err := h.playerService.GetPlayerByID(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	response := dto.PlayerResponse{
		ID:          player.ID,
		Username:    player.Username,
		Level:       player.Level,
		Experience:  player.Experience,
		Gold:        player.Gold,
		MaxDistance: player.MaxDistance,
		TotalKills:  player.TotalKills,
		CreatedAt:   player.CreatedAt,
		UpdatedAt:   player.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// UpdateMe 내 정보 수정
// @Summary      내 정보 수정
// @Description  현재 로그인한 플레이어의 정보를 수정합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.UpdateMeRequest  false  "수정할 정보"
// @Success      200   {object}  dto.PlayerResponse  "수정된 플레이어 정보"
// @Failure      401   {object}  map[string]interface{}  "인증 실패"
// @Failure      404   {object}  map[string]interface{}  "플레이어를 찾을 수 없음"
// @Router       /players/me [put]
func (h *PlayerHandler) UpdateMe(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	req := dto.UpdateMeRequest{PlayerID: playerID}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	player, err := h.playerService.GetPlayerByID(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if req.Level != nil {
		player.Level = *req.Level
	}
	if req.Experience != nil {
		player.Experience = *req.Experience
	}
	if req.Gold != nil {
		player.Gold = *req.Gold
	}
	if req.MaxDistance != nil {
		player.MaxDistance = *req.MaxDistance
	}
	if req.TotalKills != nil {
		player.TotalKills = *req.TotalKills
	}
	if err := h.playerService.UpdatePlayer(player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}

	response := dto.PlayerResponse{
		ID:          player.ID,
		Username:    player.Username,
		Level:       player.Level,
		Experience:  player.Experience,
		Gold:        player.Gold,
		MaxDistance: player.MaxDistance,
		TotalKills:  player.TotalKills,
		CreatedAt:   player.CreatedAt,
		UpdatedAt:   player.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// GetPlayers 플레이어 목록 조회
// @Summary      플레이어 목록 조회
// @Description  플레이어 목록을 페이지네이션으로 조회합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        limit   query     int  false  "조회 개수 (기본값: 10)"
// @Param        offset  query     int  false  "건너뛸 개수 (기본값: 0)"
// @Success      200     {object}  map[string][]dto.PlayerResponse  "플레이어 목록"
// @Failure      500     {object}  map[string]interface{}  "서버 오류"
// @Router       /players [get]
func (h *PlayerHandler) GetPlayers(c *gin.Context) {
	var req dto.GetPlayersRequest
	_ = c.ShouldBindQuery(&req)
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	players, err := h.playerService.GetAllPlayers(req.Limit, req.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get players"})
		return
	}
	responses := make([]dto.PlayerResponse, len(players))
	for i, p := range players {
		responses[i] = dto.PlayerResponse{
			ID:          p.ID,
			Username:    p.Username,
			Level:       p.Level,
			Experience:  p.Experience,
			Gold:        p.Gold,
			MaxDistance: p.MaxDistance,
			TotalKills:  p.TotalKills,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}
	}
	c.JSON(http.StatusOK, gin.H{"players": responses})
}

// GetPlayer 플레이어 상세 조회
// @Summary      플레이어 상세 조회
// @Description  ID로 플레이어 정보를 조회합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "플레이어 ID"
// @Success      200  {object}  dto.PlayerResponse  "플레이어 정보"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      404  {object}  map[string]interface{}  "플레이어를 찾을 수 없음"
// @Router       /players/{id} [get]
func (h *PlayerHandler) GetPlayer(c *gin.Context) {
	var pathReq dto.PlayerIDPathRequest
	if err := c.ShouldBindUri(&pathReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	player, err := h.playerService.GetPlayerByID(pathReq.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	response := dto.PlayerResponse{
		ID:          player.ID,
		Username:    player.Username,
		Level:       player.Level,
		Experience:  player.Experience,
		Gold:        player.Gold,
		MaxDistance: player.MaxDistance,
		TotalKills:  player.TotalKills,
		CreatedAt:   player.CreatedAt,
		UpdatedAt:   player.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// DeletePlayer 플레이어 삭제 (본인만)
// @Summary      플레이어 삭제
// @Description  본인 계정만 삭제할 수 있습니다 (소프트 삭제)
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "플레이어 ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      403  {object}  map[string]interface{}  "본인만 삭제 가능"
// @Failure      404  {object}  map[string]interface{}  "플레이어를 찾을 수 없음"
// @Router       /players/{id} [delete]
func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var pathReq dto.PlayerIDPathRequest
	if err := c.ShouldBindUri(&pathReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	if pathReq.ID != playerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own account"})
		return
	}

	if _, err := h.playerService.GetPlayerByID(pathReq.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if err := h.playerService.DeletePlayer(pathReq.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetLeaderboard 리더보드 조회
// @Summary      리더보드 조회
// @Description  레벨이 높은 상위 플레이어 목록을 조회합니다
// @Tags         players
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        limit   query     int  false  "조회할 플레이어 수 (기본값: 10)"
// @Success      200     {object}  map[string][]dto.PlayerResponse  "리더보드"
// @Failure      500     {object}  map[string]interface{}  "서버 오류"
// @Router       /players/leaderboard [get]
func (h *PlayerHandler) GetLeaderboard(c *gin.Context) {
	var req dto.GetLeaderboardRequest
	_ = c.ShouldBindQuery(&req)
	if req.Limit <= 0 {
		req.Limit = 10
	}

	players, err := h.playerService.GetTopPlayersByLevel(req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get leaderboard"})
		return
	}
	responses := make([]dto.PlayerResponse, len(players))
	for i, p := range players {
		responses[i] = dto.PlayerResponse{
			ID:          p.ID,
			Username:    p.Username,
			Level:       p.Level,
			Experience:  p.Experience,
			Gold:        p.Gold,
			MaxDistance: p.MaxDistance,
			TotalKills:  p.TotalKills,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}
	}
	c.JSON(http.StatusOK, gin.H{"players": responses})
}
