package repository

import (
	"errors"
	"game_eating_pizza/internal/models"
	"gorm.io/gorm"
	"sync"
)

// MockPlayerRepository는 플레이어 데이터 접근을 위한 Mock 구현체입니다
type MockPlayerRepository struct {
	players map[uint]*models.Player
	mu      sync.RWMutex
	nextID  uint
}

// NewMockPlayerRepository는 새로운 MockPlayerRepository 인스턴스를 생성합니다
func NewMockPlayerRepository() *MockPlayerRepository {
	repo := &MockPlayerRepository{
		players: make(map[uint]*models.Player),
		nextID:  1,
	}
	
	// 테스트용 초기 데이터
	repo.initTestData()
	
	return repo
}

// initTestData는 테스트용 초기 데이터를 생성합니다
func (r *MockPlayerRepository) initTestData() {
	testPlayer := &models.Player{
		ID:         1,
		Username:   "testuser",
		Level:      5,
		Gold:       1000,
		Experience: 500,
	}
	r.players[1] = testPlayer
	r.nextID = 2
}

// Create는 새로운 플레이어를 생성합니다
func (r *MockPlayerRepository) Create(player *models.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 사용자명 중복 확인
	for _, p := range r.players {
		if p.Username == player.Username {
			return errors.New("username already exists")
		}
	}

	player.ID = r.nextID
	r.nextID++
	r.players[player.ID] = player
	return nil
}

// FindByID는 ID로 플레이어를 조회합니다
func (r *MockPlayerRepository) FindByID(id uint) (*models.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	player, exists := r.players[id]
	if !exists {
		return nil, errors.New("player not found")
	}

	// 복사본 반환 (원본 수정 방지)
	result := *player
	return &result, nil
}

// FindByUsername은 사용자명으로 플레이어를 조회합니다
func (r *MockPlayerRepository) FindByUsername(username string) (*models.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, player := range r.players {
		if player.Username == username {
			// 복사본 반환
			result := *player
			return &result, nil
		}
	}

	return nil, errors.New("player not found")
}

// Update는 플레이어 정보를 업데이트합니다
func (r *MockPlayerRepository) Update(player *models.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.players[player.ID]; !exists {
		return errors.New("player not found")
	}

	r.players[player.ID] = player
	return nil
}

// UpdateGold는 플레이어의 골드를 업데이트합니다
func (r *MockPlayerRepository) UpdateGold(id uint, gold int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	player, exists := r.players[id]
	if !exists {
		return errors.New("player not found")
	}

	player.Gold = gold
	return nil
}

// FindTopPlayersByLevel은 레벨이 높은 상위 플레이어를 조회합니다
func (r *MockPlayerRepository) FindTopPlayersByLevel(limit int) ([]models.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	players := make([]models.Player, 0, len(r.players))
	for _, p := range r.players {
		players = append(players, *p)
	}

	// 레벨과 경험치로 정렬 (간단한 버블 정렬)
	for i := 0; i < len(players)-1; i++ {
		for j := i + 1; j < len(players); j++ {
			if players[i].Level < players[j].Level ||
				(players[i].Level == players[j].Level && players[i].Experience < players[j].Experience) {
				players[i], players[j] = players[j], players[i]
			}
		}
	}

	if limit > len(players) {
		limit = len(players)
	}

	return players[:limit], nil
}

// FindTopPlayersByGold은 골드가 많은 상위 플레이어를 조회합니다
func (r *MockPlayerRepository) FindTopPlayersByGold(limit int) ([]models.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	players := make([]models.Player, 0, len(r.players))
	for _, p := range r.players {
		players = append(players, *p)
	}

	// 골드로 정렬
	for i := 0; i < len(players)-1; i++ {
		for j := i + 1; j < len(players); j++ {
			if players[i].Gold < players[j].Gold {
				players[i], players[j] = players[j], players[i]
			}
		}
	}

	if limit > len(players) {
		limit = len(players)
	}

	return players[:limit], nil
}

// ExistsByUsername은 사용자명이 이미 존재하는지 확인합니다
func (r *MockPlayerRepository) ExistsByUsername(username string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, player := range r.players {
		if player.Username == username {
			return true, nil
		}
	}

	return false, nil
}

// Delete는 플레이어를 삭제합니다
func (r *MockPlayerRepository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.players[id]; !exists {
		return errors.New("player not found")
	}

	delete(r.players, id)
	return nil
}

// Transaction은 트랜잭션을 시뮬레이션합니다 (Mock에서는 단순 실행)
func (r *MockPlayerRepository) Transaction(fn func(*gorm.DB) error) error {
	// Mock에서는 트랜잭션을 시뮬레이션하지 않고 바로 실행
	return fn(nil)
}
