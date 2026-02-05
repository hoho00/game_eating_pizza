package dto

import "time"

// CreateDungeonRequest는 던전 생성 요청 DTO입니다
type CreateDungeonRequest struct {
	Name       string     `json:"name" binding:"required" example:"초급 던전"`
	Type       string     `json:"type" binding:"required" example:"normal"` // normal, event, boss
	Difficulty int        `json:"difficulty" example:"1"`
	IsActive   *bool      `json:"is_active,omitempty"` // omitempty면 기본 true
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
}

// UpdateDungeonRequest는 던전 수정 요청 DTO입니다 (부분 수정)
type UpdateDungeonRequest struct {
	Name       *string    `json:"name,omitempty"`
	Type       *string    `json:"type,omitempty"` // normal, event, boss
	Difficulty *int       `json:"difficulty,omitempty"`
	IsActive   *bool      `json:"is_active,omitempty"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
}
