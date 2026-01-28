package services

import (
	"errors"
	"game_eating_pizza/internal/models"
	"game_eating_pizza/internal/repository"
)

// WeaponService는 무기 관련 비즈니스 로직을 담당합니다
type WeaponService struct {
	weaponRepo *repository.WeaponRepository
	playerRepo *repository.PlayerRepository
}

// NewWeaponService는 새로운 WeaponService 인스턴스를 생성합니다
func NewWeaponService(
	weaponRepo *repository.WeaponRepository,
	playerRepo *repository.PlayerRepository,
) *WeaponService {
	return &WeaponService{
		weaponRepo: weaponRepo,
		playerRepo: playerRepo,
	}
}

// GetWeaponsByPlayerID는 플레이어의 무기 목록을 조회합니다
func (s *WeaponService) GetWeaponsByPlayerID(playerID uint) ([]models.Weapon, error) {
	return s.weaponRepo.FindByPlayerID(playerID)
}

// UpgradeWeapon는 무기를 강화합니다
func (s *WeaponService) UpgradeWeapon(weaponID uint) (*models.Weapon, error) {
	weapon, err := s.weaponRepo.FindByID(weaponID)
	if err != nil {
		return nil, errors.New("weapon not found")
	}

	// 강화 비용 계산 (레벨에 따라 증가)
	upgradeCost := int64(weapon.Level * 100)

	// 플레이어 골드 확인
	player, err := s.playerRepo.FindByID(weapon.PlayerID)
	if err != nil {
		return nil, errors.New("player not found")
	}

	if player.Gold < upgradeCost {
		return nil, errors.New("insufficient gold")
	}

	// 골드 차감
	player.Gold -= upgradeCost
	if err := s.playerRepo.Update(player); err != nil {
		return nil, err
	}

	// 무기 강화
	weapon.Level++
	weapon.AttackPower += 5 // 레벨당 공격력 증가
	if err := s.weaponRepo.Update(weapon); err != nil {
		return nil, err
	}

	return weapon, nil
}

// EquipWeapon는 무기를 장착합니다
func (s *WeaponService) EquipWeapon(playerID, weaponID uint) error {
	// 무기가 플레이어 소유인지 확인
	weapon, err := s.weaponRepo.FindByID(weaponID)
	if err != nil {
		return errors.New("weapon not found")
	}

	if weapon.PlayerID != playerID {
		return errors.New("weapon does not belong to player")
	}

	// 플레이어 조회 및 무기 장착
	player, err := s.playerRepo.FindByID(playerID)
	if err != nil {
		return errors.New("player not found")
	}

	player.CurrentWeaponID = &weaponID
	return s.playerRepo.Update(player)
}
