package repository

import (
	"errors"
	"game_eating_pizza/internal/models"
	"sync"
	"time"
)

// MockDungeonRepository는 던전 데이터 접근을 위한 Mock 구현체입니다
type MockDungeonRepository struct {
	dungeons map[uint]*models.Dungeon
	mu       sync.RWMutex
	nextID   uint
}

// NewMockDungeonRepository는 새로운 MockDungeonRepository 인스턴스를 생성합니다
func NewMockDungeonRepository() *MockDungeonRepository {
	repo := &MockDungeonRepository{
		dungeons: make(map[uint]*models.Dungeon),
		nextID:   1,
	}
	
	// 테스트용 초기 데이터
	repo.initTestData()
	
	return repo
}

// initTestData는 테스트용 초기 데이터를 생성합니다
func (r *MockDungeonRepository) initTestData() {
	now := time.Now()
	
	dungeon1 := &models.Dungeon{
		ID:         1,
		Name:       "Forest Dungeon",
		Type:       "normal",
		Difficulty: 1,
		IsActive:   true,
		StartTime:  &now,
	}
	r.dungeons[1] = dungeon1
	
	dungeon2 := &models.Dungeon{
		ID:         2,
		Name:       "Boss Dungeon",
		Type:       "boss",
		Difficulty: 5,
		IsActive:   true,
		StartTime:  &now,
	}
	r.dungeons[2] = dungeon2
	
	r.nextID = 3
}

// Create는 새로운 던전을 생성합니다
func (r *MockDungeonRepository) Create(dungeon *models.Dungeon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	dungeon.ID = r.nextID
	r.nextID++
	r.dungeons[dungeon.ID] = dungeon
	return nil
}

// FindByID는 ID로 던전을 조회합니다
func (r *MockDungeonRepository) FindByID(id uint) (*models.Dungeon, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dungeon, exists := r.dungeons[id]
	if !exists {
		return nil, errors.New("dungeon not found")
	}

	result := *dungeon
	return &result, nil
}

// FindActive는 활성화된 던전 목록을 조회합니다
func (r *MockDungeonRepository) FindActive() ([]models.Dungeon, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dungeons := make([]models.Dungeon, 0)
	for _, dungeon := range r.dungeons {
		if dungeon.IsActive {
			dungeons = append(dungeons, *dungeon)
		}
	}

	return dungeons, nil
}

// Update는 던전 정보를 업데이트합니다
func (r *MockDungeonRepository) Update(dungeon *models.Dungeon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.dungeons[dungeon.ID]; !exists {
		return errors.New("dungeon not found")
	}

	r.dungeons[dungeon.ID] = dungeon
	return nil
}

// Delete는 던전을 삭제합니다
func (r *MockDungeonRepository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.dungeons[id]; !exists {
		return errors.New("dungeon not found")
	}

	delete(r.dungeons, id)
	return nil
}
