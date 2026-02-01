package models

import (
	"time"
	"gorm.io/gorm"
)

// RaidSessionStatus는 레이드 세션의 상태를 나타냅니다
type RaidSessionStatus string

const (
	RaidStatusWaiting    RaidSessionStatus = "WAITING"     // 대기 중
	RaidStatusInProgress RaidSessionStatus = "IN_PROGRESS" // 진행 중
	RaidStatusCleared    RaidSessionStatus = "CLEARED"     // 클리어 완료
	RaidStatusFailed     RaidSessionStatus = "FAILED"      // 실패
)

// RaidSession은 멀티플레이 레이드 세션을 나타냅니다
// 스토리 설정: 거대 수정 거인(World Boss)을 깨우려면 여러 유저의 심장 박동을 공명시켜야 합니다
type RaidSession struct {
	ID          string            `gorm:"primaryKey;type:varchar(36)" json:"id"` // UUID
	BossID      uint              `gorm:"not null;index" json:"boss_id"`       // 보스(던전) ID
	MaxHP       uint64            `gorm:"not null" json:"max_hp"`               // 보스 최대 체력
	CurrentHP   uint64            `gorm:"not null" json:"current_hp"`         // 보스 현재 체력
	Status      RaidSessionStatus `gorm:"not null;type:varchar(20);index" json:"status"`
	StartTime   time.Time         `gorm:"not null" json:"start_time"`           // 레이드 시작 시간
	EndTime     *time.Time         `json:"end_time,omitempty"`                  // 레이드 종료 시간
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"-"`

	// 관계
	Boss        Dungeon           `gorm:"foreignKey:BossID" json:"boss,omitempty"`
	Participants []RaidParticipant `gorm:"foreignKey:RaidSessionID" json:"participants,omitempty"`
}

// TableName은 GORM이 사용할 테이블 이름을 지정합니다
func (RaidSession) TableName() string {
	return "raid_sessions"
}

// IsActive는 레이드가 진행 중인지 확인합니다
func (rs *RaidSession) IsActive() bool {
	return rs.Status == RaidStatusInProgress
}

// AddDamage는 보스에게 데미지를 입힙니다
func (rs *RaidSession) AddDamage(damage uint64) {
	if rs.CurrentHP > damage {
		rs.CurrentHP -= damage
	} else {
		rs.CurrentHP = 0
		rs.Status = RaidStatusCleared
		now := time.Now()
		rs.EndTime = &now
	}
}

// RaidParticipant는 레이드에 참여한 플레이어를 나타냅니다
type RaidParticipant struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	RaidSessionID string    `gorm:"not null;index;type:varchar(36)" json:"raid_session_id"`
	UserID        uint      `gorm:"not null;index" json:"user_id"`
	TotalDamage   uint64    `gorm:"default:0" json:"total_damage"`   // 총 기여 데미지
	StepsContributed int    `gorm:"default:0" json:"steps_contributed"` // 기여한 걸음 수 (합동 스킬용)
	JoinedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// 관계
	RaidSession RaidSession `gorm:"foreignKey:RaidSessionID" json:"raid_session,omitempty"`
	User        Player      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName은 GORM이 사용할 테이블 이름을 지정합니다
func (RaidParticipant) TableName() string {
	return "raid_participants"
}

// AddDamage는 참여자의 데미지를 추가합니다
func (rp *RaidParticipant) AddDamage(damage uint64) {
	rp.TotalDamage += damage
}

// AddSteps는 기여한 걸음 수를 추가합니다 (합동 스킬 발동 조건 확인용)
func (rp *RaidParticipant) AddSteps(steps int) {
	rp.StepsContributed += steps
}
