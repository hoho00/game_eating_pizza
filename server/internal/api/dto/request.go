package dto

import (
	"fmt"
	"strings"
	"time"
)

// FlexTime은 JSON에서 여러 시간 형식을 허용하는 타입입니다 (RFC3339, "2006-01-02T15:04:05" 등).
// T가 nil이면 omitempty로 생략된 값입니다.
type FlexTime struct {
	T *time.Time
}

// UnmarshalJSON은 "2006-01-02T15:04:05", "2006-01-02T15:04:05Z", RFC3339, "2006-01-02" 형식을 허용합니다.
func (f *FlexTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		f.T = nil
		return nil
	}
	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}
	for _, layout := range layouts {
		if parsed, err := time.Parse(layout, s); err == nil {
			f.T = &parsed
			return nil
		}
	}
	return fmt.Errorf("invalid time format: %s (use e.g. 2006-01-02T15:04:05 or 2006-01-02T15:04:05Z)", s)
}

// MarshalJSON은 RFC3339로 직렬화합니다.
func (f FlexTime) MarshalJSON() ([]byte, error) {
	if f.T == nil {
		return []byte("null"), nil
	}
	return f.T.MarshalJSON()
}

// ----- Auth -----

// RegisterRequest는 회원가입 요청 DTO입니다
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" example:"player1"`
	Password string `json:"password" binding:"required,min=6" example:"secret123"`
}

// LoginRequest는 로그인 요청 DTO입니다
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"player1"`
	Password string `json:"password" binding:"required" example:"secret123"`
}

// RefreshTokenRequest는 토큰 갱신 요청 DTO입니다 (필요 시 바디 확장)
type RefreshTokenRequest struct{}

// ----- Player -----

// GetMeRequest는 내 정보 조회 요청 DTO입니다 (PlayerID는 인증 컨텍스트에서 설정)
type GetMeRequest struct {
	PlayerID uint
}

// UpdateMeRequest는 내 정보 수정 요청 DTO입니다
type UpdateMeRequest struct {
	PlayerID    uint     // 인증 컨텍스트에서 설정
	Level       *int     `json:"level,omitempty" example:"2"`
	Experience  *int64   `json:"experience,omitempty" example:"100"`
	Gold        *int64   `json:"gold,omitempty" example:"500"`
	MaxDistance *float64 `json:"max_distance,omitempty" example:"150.5"`
	TotalKills  *int     `json:"total_kills,omitempty" example:"10"`
}

// GetLeaderboardRequest는 리더보드 조회 요청 DTO입니다
type GetLeaderboardRequest struct {
	Limit int `form:"limit"` // 기본값 10, 쿼리 파라미터
}

// GetPlayersRequest는 플레이어 목록 조회 요청 DTO입니다
type GetPlayersRequest struct {
	Limit  int `form:"limit"`  // 기본값 10
	Offset int `form:"offset"` // 기본값 0
}

// PlayerIDPathRequest는 경로 파라미터 id를 바인딩합니다 (players/:id)
type PlayerIDPathRequest struct {
	ID uint `uri:"id" binding:"required"`
}

// ----- Weapon -----

// GetWeaponsRequest는 무기 목록 조회 요청 DTO입니다
type GetWeaponsRequest struct {
	PlayerID uint
}

// CreateWeaponRequest는 무기 생성 요청 DTO입니다
type CreateWeaponRequest struct {
	PlayerID    uint    // 인증 컨텍스트에서 설정
	Name        string  `json:"name" binding:"required" example:"Iron Sword"`
	Type        string  `json:"type" binding:"required" example:"sword"` // sword, bow, staff
	AttackPower int     `json:"attack_power" example:"10"`
	AttackSpeed float64 `json:"attack_speed" example:"1.0"`
	Rarity      string  `json:"rarity" example:"common"` // common, rare, epic, legendary
}

// WeaponIDPathRequest는 경로 파라미터 id를 바인딩합니다 (weapons/:id)
type WeaponIDPathRequest struct {
	ID uint `uri:"id" binding:"required"`
}

// UpdateWeaponRequest는 무기 수정 요청 DTO입니다 (부분 수정)
type UpdateWeaponRequest struct {
	Name        *string  `json:"name,omitempty" example:"Steel Sword"`
	Type        *string  `json:"type,omitempty" example:"sword"`
	AttackPower *int     `json:"attack_power,omitempty" example:"15"`
	AttackSpeed *float64 `json:"attack_speed,omitempty" example:"1.2"`
	Rarity      *string  `json:"rarity,omitempty" example:"rare"`
	Level       *int     `json:"level,omitempty" example:"2"`
}

// UpgradeWeaponRequest는 무기 강화 요청 DTO입니다
type UpgradeWeaponRequest struct {
	PlayerID uint // 인증 컨텍스트에서 설정
	WeaponID uint // 경로에서 설정
}

// EquipWeaponRequest는 무기 장착 요청 DTO입니다
type EquipWeaponRequest struct {
	PlayerID uint
	WeaponID uint
}

// ----- Dungeon -----

// CreateDungeonRequest는 던전 생성 요청 DTO입니다
type CreateDungeonRequest struct {
	Name       string   `json:"name" binding:"required" example:"초급 던전"`
	Type       string   `json:"type" binding:"required" example:"normal"` // normal, event, boss
	Difficulty int      `json:"difficulty" example:"1"`
	IsActive   *bool    `json:"is_active,omitempty"` // omitempty면 기본 true
	StartTime  FlexTime `json:"start_time,omitempty"`
	EndTime    FlexTime `json:"end_time,omitempty"`
}

// UpdateDungeonRequest는 던전 수정 요청 DTO입니다 (부분 수정)
type UpdateDungeonRequest struct {
	Name       *string   `json:"name,omitempty" example:"수정된 던전"`
	Type       *string   `json:"type,omitempty" example:"event"`
	Difficulty *int      `json:"difficulty,omitempty" example:"2"`
	IsActive   *bool     `json:"is_active,omitempty" example:"true"`
	StartTime  *FlexTime `json:"start_time,omitempty"`
	EndTime    *FlexTime `json:"end_time,omitempty"`
}

// DungeonIDPathRequest는 경로 파라미터 id를 바인딩합니다 (dungeons/:id)
type DungeonIDPathRequest struct {
	ID uint `uri:"id" binding:"required"`
}

// GetDungeonRequest는 던전 상세 조회 요청 DTO입니다
type GetDungeonRequest struct {
	DungeonID uint
}

// DeleteDungeonRequest는 던전 삭제 요청 DTO입니다
type DeleteDungeonRequest struct {
	DungeonID uint
}

// EnterDungeonRequest는 던전 입장 요청 DTO입니다
type EnterDungeonRequest struct {
	PlayerID  uint
	DungeonID uint
}

// ClearDungeonRequest는 던전 클리어 요청 DTO입니다
type ClearDungeonRequest struct {
	PlayerID  uint
	DungeonID uint
}
