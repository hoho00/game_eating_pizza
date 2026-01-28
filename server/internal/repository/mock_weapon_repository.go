package repository

import (
	"errors"
	"game_eating_pizza/internal/models"
	"sync"
)

// MockWeaponRepository는 무기 데이터 접근을 위한 Mock 구현체입니다
type MockWeaponRepository struct {
	weapons map[uint]*models.Weapon
	mu      sync.RWMutex
	nextID  uint
}

// NewMockWeaponRepository는 새로운 MockWeaponRepository 인스턴스를 생성합니다
func NewMockWeaponRepository() *MockWeaponRepository {
	repo := &MockWeaponRepository{
		weapons: make(map[uint]*models.Weapon),
		nextID:  1,
	}
	
	// 테스트용 초기 데이터
	repo.initTestData()
	
	return repo
}

// initTestData는 테스트용 초기 데이터를 생성합니다
func (r *MockWeaponRepository) initTestData() {
	testWeapon := &models.Weapon{
		ID:          1,
		PlayerID:    1,
		Name:        "Basic Sword",
		Type:        "sword",
		AttackPower: 10,
		AttackSpeed: 1.0,
		Rarity:      "common",
		Level:       1,
	}
	r.weapons[1] = testWeapon
	r.nextID = 2
}

// Create는 새로운 무기를 생성합니다
func (r *MockWeaponRepository) Create(weapon *models.Weapon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	weapon.ID = r.nextID
	r.nextID++
	r.weapons[weapon.ID] = weapon
	return nil
}

// FindByID는 ID로 무기를 조회합니다
func (r *MockWeaponRepository) FindByID(id uint) (*models.Weapon, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	weapon, exists := r.weapons[id]
	if !exists {
		return nil, errors.New("weapon not found")
	}

	result := *weapon
	return &result, nil
}

// FindByPlayerID는 플레이어 ID로 무기 목록을 조회합니다
func (r *MockWeaponRepository) FindByPlayerID(playerID uint) ([]models.Weapon, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	weapons := make([]models.Weapon, 0)
	for _, weapon := range r.weapons {
		if weapon.PlayerID == playerID {
			weapons = append(weapons, *weapon)
		}
	}

	return weapons, nil
}

// Update는 무기 정보를 업데이트합니다
func (r *MockWeaponRepository) Update(weapon *models.Weapon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.weapons[weapon.ID]; !exists {
		return errors.New("weapon not found")
	}

	r.weapons[weapon.ID] = weapon
	return nil
}

// Delete는 무기를 삭제합니다
func (r *MockWeaponRepository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.weapons[id]; !exists {
		return errors.New("weapon not found")
	}

	delete(r.weapons, id)
	return nil
}
