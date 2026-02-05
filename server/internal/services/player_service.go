package services

import (
	"game_eating_pizza/internal/models"
	"game_eating_pizza/internal/repository"
)

// PlayerService는 플레이어 관련 비즈니스 로직을 담당합니다
type PlayerService struct {
	playerRepo repository.PlayerRepositoryInterface
	weaponRepo repository.WeaponRepositoryInterface
}

// NewPlayerService는 새로운 PlayerService 인스턴스를 생성합니다
func NewPlayerService(
	playerRepo repository.PlayerRepositoryInterface,
	weaponRepo repository.WeaponRepositoryInterface,
) *PlayerService {
	return &PlayerService{
		playerRepo: playerRepo,
		weaponRepo: weaponRepo,
	}
}

// GetPlayerByID는 ID로 플레이어를 조회합니다
func (s *PlayerService) GetPlayerByID(id uint) (*models.Player, error) {
	return s.playerRepo.FindByID(id)
}

// GetTopPlayersByLevel은 레벨이 높은 상위 플레이어를 조회합니다
func (s *PlayerService) GetTopPlayersByLevel(limit int) ([]models.Player, error) {
	return s.playerRepo.FindTopPlayersByLevel(limit)
}

// GetAllPlayers는 플레이어 목록을 페이지네이션으로 조회합니다
func (s *PlayerService) GetAllPlayers(limit, offset int) ([]models.Player, error) {
	return s.playerRepo.FindAll(limit, offset)
}

// UpdatePlayer는 플레이어 정보를 수정합니다
func (s *PlayerService) UpdatePlayer(player *models.Player) error {
	return s.playerRepo.Update(player)
}

// DeletePlayer는 플레이어를 삭제합니다 (소프트 삭제)
func (s *PlayerService) DeletePlayer(id uint) error {
	return s.playerRepo.Delete(id)
}

// CreatePlayer는 새로운 플레이어를 생성하고 기본 무기를 제공합니다
func (s *PlayerService) CreatePlayer(username string) (*models.Player, error) {
	player := &models.Player{
		Username: username,
		Level:    1,
		Gold:     0,
	}

	if err := s.playerRepo.Create(player); err != nil {
		return nil, err
	}

	// 기본 무기 생성
	defaultWeapon := &models.Weapon{
		PlayerID:    player.ID,
		Name:        "Basic Sword",
		Type:        "sword",
		AttackPower: 10,
		AttackSpeed: 1.0,
		Rarity:      "common",
		Level:       1,
	}

	if err := s.weaponRepo.Create(defaultWeapon); err != nil {
		return nil, err
	}

	player.CurrentWeaponID = &defaultWeapon.ID
	s.playerRepo.Update(player)

	return player, nil
}
