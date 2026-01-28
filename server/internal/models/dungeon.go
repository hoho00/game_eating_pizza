package models

import (
	"time"
)

// Dungeon는 던전 정보를 나타냅니다
type Dungeon struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"not null;size:100" json:"name"`
	Type       string     `gorm:"not null;size:20" json:"type"` // normal, event, boss
	Difficulty int        `gorm:"default:1;index" json:"difficulty"`
	IsActive   bool       `gorm:"default:true;index" json:"is_active"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// TableName은 GORM이 사용할 테이블 이름을 지정합니다
func (Dungeon) TableName() string {
	return "dungeons"
}
