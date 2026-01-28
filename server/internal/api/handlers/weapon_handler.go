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

// GetWeapons는 플레이어의 무기 목록을 조회합니다
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

// CreateWeapon는 새로운 무기를 생성합니다
func (h *WeaponHandler) CreateWeapon(c *gin.Context) {
	// TODO: 구현 필요
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented yet",
	})
}

// UpgradeWeapon는 무기를 강화합니다
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

// EquipWeapon는 무기를 장착합니다
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
