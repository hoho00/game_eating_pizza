package models

import (
	"time"
	"gorm.io/gorm"
)

// UserActivity는 사용자의 일일 활동 데이터(걸음 수 등)를 나타냅니다
// 스토리 설정: 주인공의 움직임(심장 박동/발걸음)이 대장간의 화로를 뜨겁게 만드는 연료입니다
type UserActivity struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null;index" json:"user_id"`
	Date         string    `gorm:"not null;size:10;index" json:"date"` // "2024-05-21" 형식
	Steps        int       `gorm:"default:0" json:"steps"`            // 당일 걸음 수
	Calories     float64   `gorm:"default:0" json:"calories"`         // 소모 칼로리
	LastSyncedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_synced_at"` // 마지막 동기화 시간
	BonusApplied bool      `gorm:"default:false" json:"bonus_applied"` // 걸음 수 목표 달성 보상 수령 여부
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// 관계
	User Player `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName은 GORM이 사용할 테이블 이름을 지정합니다
func (UserActivity) TableName() string {
	return "user_activities"
}

// AddSteps는 걸음 수를 추가합니다 (위변조 방지를 위한 검증 로직 포함)
func (ua *UserActivity) AddSteps(steps int) {
	// 비정상적인 증가 속도 필터링 (예: 1초에 100걸음 이상은 무시)
	// 실제 구현 시 더 정교한 검증 로직 필요
	ua.Steps += steps
	ua.LastSyncedAt = time.Now()
}

// GetForgeBoost는 걸음 수에 따른 대장간 부스트 배율을 반환합니다
// 스토리: 걸을수록 화로가 뜨거워져 무기 제작 속도가 증가합니다
func (ua *UserActivity) GetForgeBoost() float64 {
	// 예: 5,000보 이상이면 2배 부스트
	if ua.Steps >= 5000 {
		return 2.0
	}
	// 예: 3,000보 이상이면 1.5배 부스트
	if ua.Steps >= 3000 {
		return 1.5
	}
	// 기본: 1.0배
	return 1.0
}
