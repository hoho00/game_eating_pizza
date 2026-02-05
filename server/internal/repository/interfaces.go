package repository

import (
	"game_eating_pizza/internal/models"
	"gorm.io/gorm"
)

// PlayerRepositoryInterface는 플레이어 데이터 접근 인터페이스입니다
type PlayerRepositoryInterface interface {
	Create(player *models.Player) error
	FindByID(id uint) (*models.Player, error)
	FindByUsername(username string) (*models.Player, error)
	Update(player *models.Player) error
	UpdateGold(id uint, gold int64) error
	FindAll(limit, offset int) ([]models.Player, error)
	FindTopPlayersByLevel(limit int) ([]models.Player, error)
	FindTopPlayersByGold(limit int) ([]models.Player, error)
	ExistsByUsername(username string) (bool, error)
	Delete(id uint) error
	Transaction(fn func(*gorm.DB) error) error
}

// WeaponRepositoryInterface는 무기 데이터 접근 인터페이스입니다
type WeaponRepositoryInterface interface {
	Create(weapon *models.Weapon) error
	FindByID(id uint) (*models.Weapon, error)
	FindByPlayerID(playerID uint) ([]models.Weapon, error)
	Update(weapon *models.Weapon) error
	Delete(id uint) error
}

// DungeonRepositoryInterface는 던전 데이터 접근 인터페이스입니다
type DungeonRepositoryInterface interface {
	Create(dungeon *models.Dungeon) error
	FindByID(id uint) (*models.Dungeon, error)
	FindAll() ([]models.Dungeon, error)
	FindActive() ([]models.Dungeon, error)
	Update(dungeon *models.Dungeon) error
	Delete(id uint) error
}
