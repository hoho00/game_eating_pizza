package handlers

import (
	"game_eating_pizza/internal/services"
	"net/http"
	"strconv"

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

	weapons, err := h.weaponService.GetWeaponsByPlayerID(uint(playerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get weapons",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"weapons": weapons,
	})
}

// CreateWeapon 무기 생성
// @Summary      무기 생성
// @Description  새로운 무기를 생성합니다
// @Tags         weapons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      201      {object}  map[string]interface{}  "무기 생성 성공"
// @Failure      401      {object}  map[string]interface{}  "인증 실패"
// @Router       /weapons [post]
func (h *WeaponHandler) CreateWeapon(c *gin.Context) {
	// TODO: 구현 필요
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}

// UpgradeWeapon 무기 강화
// @Summary      무기 강화
// @Description  무기를 강화하여 공격력을 증가시킵니다
// @Tags         weapons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "무기 ID"
// @Success      200  {object}  models.Weapon  "강화된 무기 정보"
// @Failure      400  {object}  map[string]interface{}  "잘못된 요청"
// @Failure      401  {object}  map[string]interface{}  "인증 실패"
// @Failure      500  {object}  map[string]interface{}  "서버 오류"
// @Router       /weapons/{id}/upgrade [put]
func (h *WeaponHandler) UpgradeWeapon(c *gin.Context) {
	weaponIDStr := c.Param("id")
	weaponID, err := strconv.ParseUint(weaponIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid weapon ID",
		})
		return
	}

	weapon, err := h.weaponService.UpgradeWeapon(uint(weaponID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upgrade weapon",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, weapon)
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
	weaponIDStr := c.Param("id")
	weaponID, err := strconv.ParseUint(weaponIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid weapon ID",
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

	err = h.weaponService.EquipWeapon(uint(playerID), uint(weaponID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to equip weapon",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Weapon equipped successfully",
	})
}
