package repository

import (
	"game_eating_pizza/internal/config"
	"gorm.io/gorm"
)

// Repositories는 모든 Repository를 담는 구조체입니다
type Repositories struct {
	Player  PlayerRepositoryInterface
	Weapon  WeaponRepositoryInterface
	Dungeon DungeonRepositoryInterface
}

// NewRepositories는 실제 데이터베이스 Repository를 생성합니다
func NewRepositories(db *gorm.DB, cfg *config.Config) *Repositories {
	return &Repositories{
		Player:  NewPlayerRepository(db),
		Weapon:  NewWeaponRepository(db),
		Dungeon: NewDungeonRepository(db),
	}
}
