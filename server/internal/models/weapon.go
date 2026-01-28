package models

import (
	"time"
)

// Weapon는 무기 정보를 나타냅니다
type Weapon struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PlayerID    uint      `gorm:"not null;index" json:"player_id"`
	Name        string    `gorm:"not null;size:100" json:"name"`
	Type        string    `gorm:"not null;size:20" json:"type"` // sword, bow, staff
	AttackPower int       `gorm:"default:10" json:"attack_power"`
	AttackSpeed float64   `gorm:"default:1.0" json:"attack_speed"`
	Rarity      string    `gorm:"default:common;size:20" json:"rarity"` // common, rare, epic, legendary
	Level       int       `gorm:"default:1" json:"level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// 관계
	Player Player `gorm:"foreignKey:PlayerID" json:"-"`
}

// TableName은 GORM이 사용할 테이블 이름을 지정합니다
func (Weapon) TableName() string {
	return "weapons"
}
