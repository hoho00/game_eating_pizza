package handlers

import (
	"game_eating_pizza/internal/api/params"
	"game_eating_pizza/internal/api/dto"
	"game_eating_pizza/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WeaponHandler는 무기 관련 핸들러입니다
type WeaponHandler struct {
	weaponService *services.WeaponService
}

// NewWeaponHandler는 새로운 WeaponHandler를 생성합니다
func NewWeaponHandler(weaponService *services.WeaponService) *WeaponHandler {
	return &WeaponHandler{
		weaponService: weaponService,
	}
}

// GetWeapons 무기 목록 조회
// @Summary      무기 목록 조회
// @Description  현재 로그인한 플레이어의 무기 목록을 조회합니다
// @Tags         weapons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  map[string]interface{}  "무기 목록"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Failure      500      {object}  map[string]interface{}  "서버 오류"
// @Router       /weapons [get]
func (h *WeaponHandler) GetWeapons(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	req := dto.GetWeaponsRequest{PlayerID: playerID}

	weapons, err := h.weaponService.GetWeaponsByPlayerID(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get weapons"})
		return
	}

	responses := make([]dto.WeaponResponse, len(weapons))
	for i, weapon := range weapons {
		responses[i] = dto.WeaponResponse{
			ID:          weapon.ID,
			PlayerID:    weapon.PlayerID,
			Name:        weapon.Name,
			Type:        weapon.Type,
			AttackPower: weapon.AttackPower,
			AttackSpeed: weapon.AttackSpeed,
			Rarity:      weapon.Rarity,
			Level:       weapon.Level,
			CreatedAt:   weapon.CreatedAt,
			UpdatedAt:   weapon.UpdatedAt,
		}
	}
	c.JSON(http.StatusOK, gin.H{"weapons": responses})
}

// CreateWeapon 무기 생성
// @Summary      무기 생성
// @Description  새로운 무기를 생성합니다
// @Tags         weapons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.CreateWeaponRequest  true  "무기 생성 정보"
// @Success      201   {object}  map[string]interface{}  "무기 생성 성공"
// @Failure      401   {object}  map[string]interface{}  "인증 실패"
// @Router       /weapons [post]
func (h *WeaponHandler) CreateWeapon(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var req dto.CreateWeaponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}
	req.PlayerID = playerID

	// TODO: weaponService.CreateWeapon(req) 구현
	_ = req
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

// UpgradeWeapon 무기 강화
// @Summary      무기 강화
// @Description  무기를 강화하여 공격력을 증가시킵니다
// @Tags         weapons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "무기 ID"
// @Success      200  {object}  dto.WeaponResponse  "강화된 무기 정보"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      500  {object}  map[string]interface{}  "서버 오류"
// @Router       /weapons/{id}/upgrade [put]
func (h *WeaponHandler) UpgradeWeapon(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var pathReq dto.WeaponIDPathRequest
	if err := c.ShouldBindUri(&pathReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid weapon ID"})
		return
	}
	req := dto.UpgradeWeaponRequest{PlayerID: playerID, WeaponID: pathReq.ID}

	weapon, err := h.weaponService.UpgradeWeapon(req.WeaponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade weapon", "details": err.Error()})
		return
	}

	response := dto.WeaponResponse{
		ID:          weapon.ID,
		PlayerID:    weapon.PlayerID,
		Name:        weapon.Name,
		Type:        weapon.Type,
		AttackPower: weapon.AttackPower,
		AttackSpeed: weapon.AttackSpeed,
		Rarity:      weapon.Rarity,
		Level:       weapon.Level,
		CreatedAt:   weapon.CreatedAt,
		UpdatedAt:   weapon.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// EquipWeapon 무기 장착
// @Summary      무기 장착
// @Description  무기를 장착하여 현재 무기로 설정합니다
// @Tags         weapons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "무기 ID"
// @Success      200  {object}  map[string]interface{}  "장착 성공"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      500  {object}  map[string]interface{}  "서버 오류"
// @Router       /weapons/{id}/equip [put]
func (h *WeaponHandler) EquipWeapon(c *gin.Context) {
	playerID, ok := params.GetAuthenticatedPlayerID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var pathReq dto.WeaponIDPathRequest
	if err := c.ShouldBindUri(&pathReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid weapon ID"})
		return
	}
	req := dto.EquipWeaponRequest{PlayerID: playerID, WeaponID: pathReq.ID}

	err := h.weaponService.EquipWeapon(req.PlayerID, req.WeaponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to equip weapon", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Weapon equipped successfully"})
}
