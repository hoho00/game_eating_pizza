package repository

import (
	"game_eating_pizza/internal/models"
	"gorm.io/gorm"
)

// WeaponRepository는 무기 데이터 접근을 담당합니다
type WeaponRepository struct {
	db *gorm.DB
}

// NewWeaponRepository는 새로운 WeaponRepository 인스턴스를 생성합니다
func NewWeaponRepository(db *gorm.DB) *WeaponRepository {
	return &WeaponRepository{db: db}
}

// Create는 새로운 무기를 생성합니다
func (r *WeaponRepository) Create(weapon *models.Weapon) error {
	return r.db.Create(weapon).Error
}

// FindByID는 ID로 무기를 조회합니다
func (r *WeaponRepository) FindByID(id uint) (*models.Weapon, error) {
	var weapon models.Weapon
	err := r.db.First(&weapon, id).Error
	if err != nil {
		return nil, err
	}
	return &weapon, nil
}

// FindByPlayerID는 플레이어 ID로 무기 목록을 조회합니다
func (r *WeaponRepository) FindByPlayerID(playerID uint) ([]models.Weapon, error) {
	var weapons []models.Weapon
	err := r.db.Where("player_id = ?", playerID).Find(&weapons).Error
	return weapons, err
}

// Update는 무기 정보를 업데이트합니다
func (r *WeaponRepository) Update(weapon *models.Weapon) error {
	return r.db.Save(weapon).Error
}

// Delete는 무기를 삭제합니다
func (r *WeaponRepository) Delete(id uint) error {
	return r.db.Delete(&models.Weapon{}, id).Error
}
