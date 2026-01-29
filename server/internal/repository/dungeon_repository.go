package repository

import (
	"game_eating_pizza/internal/models"
	"gorm.io/gorm"
)

// DungeonRepository는 던전 데이터 접근을 담당합니다
// DungeonRepositoryInterface를 구현합니다
type DungeonRepository struct {
	db *gorm.DB
}

// DungeonRepository가 인터페이스를 구현하는지 컴파일 타임에 확인
var _ DungeonRepositoryInterface = (*DungeonRepository)(nil)

// NewDungeonRepository는 새로운 DungeonRepository 인스턴스를 생성합니다
func NewDungeonRepository(db *gorm.DB) *DungeonRepository {
	return &DungeonRepository{db: db}
}

// Create는 새로운 던전을 생성합니다
func (r *DungeonRepository) Create(dungeon *models.Dungeon) error {
	return r.db.Create(dungeon).Error
}

// FindByID는 ID로 던전을 조회합니다
func (r *DungeonRepository) FindByID(id uint) (*models.Dungeon, error) {
	var dungeon models.Dungeon
	err := r.db.First(&dungeon, id).Error
	if err != nil {
		return nil, err
	}
	return &dungeon, nil
}

// FindAll은 전체 던전 목록을 조회합니다
func (r *DungeonRepository) FindAll() ([]models.Dungeon, error) {
	var dungeons []models.Dungeon
	err := r.db.Find(&dungeons).Error
	return dungeons, err
}

// FindActive는 활성화된 던전 목록을 조회합니다
func (r *DungeonRepository) FindActive() ([]models.Dungeon, error) {
	var dungeons []models.Dungeon
	err := r.db.Where("is_active = ?", true).Find(&dungeons).Error
	return dungeons, err
}

// Update는 던전 정보를 업데이트합니다
func (r *DungeonRepository) Update(dungeon *models.Dungeon) error {
	return r.db.Save(dungeon).Error
}

// Delete는 던전을 삭제합니다
func (r *DungeonRepository) Delete(id uint) error {
	return r.db.Delete(&models.Dungeon{}, id).Error
}
