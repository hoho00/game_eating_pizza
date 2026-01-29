package services

import (
	"game_eating_pizza/internal/models"
	"game_eating_pizza/internal/repository"
)

// DungeonService는 던전 관련 비즈니스 로직을 담당합니다
type DungeonService struct {
	dungeonRepo repository.DungeonRepositoryInterface
}

// NewDungeonService는 새로운 DungeonService 인스턴스를 생성합니다
func NewDungeonService(dungeonRepo repository.DungeonRepositoryInterface) *DungeonService {
	return &DungeonService{
		dungeonRepo: dungeonRepo,
	}
}

// GetDungeonByID는 ID로 던전을 조회합니다
func (s *DungeonService) GetDungeonByID(id uint) (*models.Dungeon, error) {
	return s.dungeonRepo.FindByID(id)
}

// GetAllDungeons는 전체 던전 목록을 조회합니다
func (s *DungeonService) GetAllDungeons() ([]models.Dungeon, error) {
	return s.dungeonRepo.FindAll()
}

// GetActiveDungeons는 활성화된 던전 목록을 조회합니다
func (s *DungeonService) GetActiveDungeons() ([]models.Dungeon, error) {
	return s.dungeonRepo.FindActive()
}

// EnterDungeon는 던전에 입장합니다
func (s *DungeonService) EnterDungeon(playerID, dungeonID uint) error {
	// TODO: 던전 입장 로직 구현
	// - 던전 존재 확인
	// - 플레이어 레벨/요구사항 확인
	// - 던전 입장 기록 생성
	return nil
}
