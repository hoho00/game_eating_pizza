package dto

import "time"

// PlayerResponse는 플레이어 정보 응답 DTO입니다
type PlayerResponse struct {
	ID          uint    `json:"id"`
	Username    string  `json:"username"`
	Level       int     `json:"level"`
	Experience  int64   `json:"experience"`
	Gold        int64   `json:"gold"`
	MaxDistance float64 `json:"max_distance"`
	TotalKills  int     `json:"total_kills"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// WeaponResponse는 무기 정보 응답 DTO입니다
type WeaponResponse struct {
	ID          uint    `json:"id"`
	PlayerID    uint    `json:"player_id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	AttackPower int     `json:"attack_power"`
	AttackSpeed float64 `json:"attack_speed"`
	Rarity      string  `json:"rarity"`
	Level       int     `json:"level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DungeonResponse는 던전 정보 응답 DTO입니다
type DungeonResponse struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Difficulty int        `json:"difficulty"`
	IsActive   bool       `json:"is_active"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
