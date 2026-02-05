package repository

import (
	"game_eating_pizza/internal/models"
	"gorm.io/gorm"
)

// PlayerRepository는 플레이어 데이터 접근을 담당합니다 (JPA Repository 패턴)
// PlayerRepositoryInterface를 구현합니다
type PlayerRepository struct {
	db *gorm.DB
}

// PlayerRepository가 인터페이스를 구현하는지 컴파일 타임에 확인
var _ PlayerRepositoryInterface = (*PlayerRepository)(nil)

// NewPlayerRepository는 새로운 PlayerRepository 인스턴스를 생성합니다
func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// Create는 새로운 플레이어를 생성합니다
func (r *PlayerRepository) Create(player *models.Player) error {
	return r.db.Create(player).Error
}

// FindByID는 ID로 플레이어를 조회합니다 (관계 포함)
func (r *PlayerRepository) FindByID(id uint) (*models.Player, error) {
	var player models.Player
	err := r.db.
		Preload("Weapons").
		Preload("CurrentWeapon").
		First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// FindByUsername은 사용자명으로 플레이어를 조회합니다
func (r *PlayerRepository) FindByUsername(username string) (*models.Player, error) {
	var player models.Player
	err := r.db.Where("username = ?", username).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// Update는 플레이어 정보를 업데이트합니다
func (r *PlayerRepository) Update(player *models.Player) error {
	return r.db.Save(player).Error
}

// UpdateGold는 플레이어의 골드를 업데이트합니다 (선택적 업데이트)
func (r *PlayerRepository) UpdateGold(id uint, gold int64) error {
	return r.db.Model(&models.Player{}).
		Where("id = ?", id).
		Update("gold", gold).Error
}

// FindAll은 플레이어 목록을 페이지네이션으로 조회합니다
func (r *PlayerRepository) FindAll(limit, offset int) ([]models.Player, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	var players []models.Player
	err := r.db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&players).Error
	return players, err
}

// FindTopPlayersByLevel은 레벨이 높은 상위 플레이어를 조회합니다
func (r *PlayerRepository) FindTopPlayersByLevel(limit int) ([]models.Player, error) {
	var players []models.Player
	err := r.db.
		Order("level DESC, experience DESC").
		Limit(limit).
		Find(&players).Error
	return players, err
}

// FindTopPlayersByGold은 골드가 많은 상위 플레이어를 조회합니다
func (r *PlayerRepository) FindTopPlayersByGold(limit int) ([]models.Player, error) {
	var players []models.Player
	err := r.db.
		Order("gold DESC").
		Limit(limit).
		Find(&players).Error
	return players, err
}

// ExistsByUsername은 사용자명이 이미 존재하는지 확인합니다
func (r *PlayerRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Player{}).
		Where("username = ?", username).
		Count(&count).Error
	return count > 0, err
}

// Delete는 플레이어를 소프트 삭제합니다 (DeletedAt 설정)
func (r *PlayerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Player{}, id).Error
}

// Transaction은 트랜잭션 내에서 함수를 실행합니다
func (r *PlayerRepository) Transaction(fn func(*gorm.DB) error) error {
	return r.db.Transaction(fn)
}
