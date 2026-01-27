# Go ORM 추천 가이드

## 개요

SpringBoot + JPA 경험이 좋았던 개발자를 위한 Go ORM 추천 가이드입니다.

---

## 추천: GORM (1순위)

### 왜 GORM인가?

**JPA와 가장 유사한 개발 경험**을 제공하는 Go ORM입니다.

#### 주요 특징
- ✅ **풍부한 기능**: 관계 설정, 마이그레이션, 훅, 트랜잭션 등
- ✅ **직관적인 API**: JPA와 유사한 메서드 체이닝
- ✅ **활발한 커뮤니티**: 가장 인기 있는 Go ORM (GitHub 38k+ stars)
- ✅ **우수한 문서**: 한글 문서도 제공
- ✅ **다양한 DB 지원**: PostgreSQL, MySQL, SQLite, SQL Server 등

#### JPA와의 유사점
```go
// JPA 스타일
type Player struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex;not null"`
    Level     int       `gorm:"default:1"`
    Gold      int64     `gorm:"default:0"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Weapons   []Weapon  `gorm:"foreignKey:PlayerID"`
}

// Repository 패턴 (JPA Repository와 유사)
type PlayerRepository struct {
    db *gorm.DB
}

func (r *PlayerRepository) FindByID(id uint) (*Player, error) {
    var player Player
    err := r.db.Preload("Weapons").First(&player, id).Error
    return &player, err
}
```

### 설치
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres  # PostgreSQL 사용 시
# 또는
go get -u gorm.io/driver/mysql     # MySQL 사용 시
```

### 기본 사용 예시
```go
package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// 모델 정의
type Player struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex;not null"`
    Level     int       `gorm:"default:1"`
    Gold      int64     `gorm:"default:0"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 데이터베이스 연결
func ConnectDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=postgres dbname=game_db port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // 마이그레이션 (JPA의 @Entity와 유사)
    db.AutoMigrate(&Player{})
    
    return db, nil
}

// CRUD 작업
func CreatePlayer(db *gorm.DB, username string) (*Player, error) {
    player := &Player{
        Username: username,
        Level:    1,
        Gold:     0,
    }
    result := db.Create(player)
    return player, result.Error
}

func FindPlayerByID(db *gorm.DB, id uint) (*Player, error) {
    var player Player
    result := db.First(&player, id)
    return &player, result.Error
}

func UpdatePlayerGold(db *gorm.DB, id uint, gold int64) error {
    return db.Model(&Player{}).Where("id = ?", id).Update("gold", gold).Error
}
```

---

## 대안: Ent (2순위)

### 특징
- **타입 안전성**: 코드 생성 기반으로 컴파일 타임 타입 체크
- **그래프 기반**: 복잡한 관계를 쉽게 표현
- **Facebook 개발**: 메타(Facebook)에서 개발 및 유지보수

### 언제 사용하면 좋은가?
- 복잡한 데이터 관계가 많은 경우
- 타입 안전성이 매우 중요한 경우
- 코드 생성 기반이 괜찮은 경우

### 단점
- GORM보다 학습 곡선이 가파름
- JPA와의 유사도가 낮음

---

## 대안: sqlc (성능 중심)

### 특징
- **SQL 직접 작성**: SQL을 직접 작성하고 타입 안전한 Go 코드 생성
- **높은 성능**: ORM 오버헤드 없음
- **명시적 쿼리**: 모든 쿼리를 명시적으로 관리

### 언제 사용하면 좋은가?
- 성능이 매우 중요한 경우 (게임 서버)
- SQL을 직접 작성하는 것을 선호하는 경우
- 복잡한 쿼리가 많은 경우

### 단점
- JPA와의 유사도가 매우 낮음
- SQL을 직접 작성해야 함

---

## 게임 프로젝트 추천: GORM

### 이유
1. **빠른 개발**: JPA 경험으로 빠르게 적응 가능
2. **관계 관리**: 플레이어-무기-던전 관계를 쉽게 표현
3. **마이그레이션**: 스키마 변경을 쉽게 관리
4. **충분한 성능**: 방치형 게임의 트래픽에는 충분

### 성능 최적화 팁
```go
// 1. Preload로 N+1 문제 해결 (JPA의 @EntityGraph와 유사)
db.Preload("Weapons").Preload("Weapons.SpecialEffects").Find(&players)

// 2. Select로 필요한 컬럼만 조회
db.Select("id", "username", "level").Find(&players)

// 3. 인덱스 활용
type Player struct {
    Username string `gorm:"uniqueIndex:idx_username"`
    Level    int    `gorm:"index:idx_level"`
}

// 4. 배치 작업
db.CreateInBatches(players, 100)

// 5. 트랜잭션 사용
db.Transaction(func(tx *gorm.DB) error {
    // 여러 작업을 하나의 트랜잭션으로
    return nil
})
```

---

## 프로젝트 적용 예시

### 모델 정의 (게임 데이터)
```go
// internal/models/player.go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Player struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex;not null;size:50"`
    Level     int       `gorm:"default:1;index"`
    Experience int64    `gorm:"default:0"`
    Gold      int64     `gorm:"default:0;index"`
    MaxDistance float64 `gorm:"default:0"`
    TotalKills int      `gorm:"default:0"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"`
    
    // 관계
    Weapons    []Weapon       `gorm:"foreignKey:PlayerID"`
    CurrentWeaponID *uint
    CurrentWeapon   *Weapon   `gorm:"foreignKey:CurrentWeaponID"`
}

// internal/models/weapon.go
type Weapon struct {
    ID           uint      `gorm:"primaryKey"`
    PlayerID     uint      `gorm:"not null;index"`
    Name         string    `gorm:"not null;size:100"`
    Type         string    `gorm:"not null;size:20"` // sword, bow, staff
    AttackPower  int       `gorm:"default:10"`
    AttackSpeed  float64   `gorm:"default:1.0"`
    Rarity       string    `gorm:"default:common;size:20"` // common, rare, epic, legendary
    Level        int       `gorm:"default:1"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    
    // 관계
    Player       Player    `gorm:"foreignKey:PlayerID"`
}

// internal/models/dungeon.go
type Dungeon struct {
    ID            uint      `gorm:"primaryKey"`
    Name          string    `gorm:"not null;size:100"`
    Type          string    `gorm:"not null;size:20"` // normal, event, boss
    Difficulty    int       `gorm:"default:1;index"`
    IsActive      bool      `gorm:"default:true;index"`
    StartTime     *time.Time
    EndTime       *time.Time
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
```

### Repository 패턴 (JPA Repository 스타일)
```go
// internal/repository/player_repository.go
package repository

import (
    "game_eating_pizza/internal/models"
    "gorm.io/gorm"
)

type PlayerRepository struct {
    db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
    return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(player *models.Player) error {
    return r.db.Create(player).Error
}

func (r *PlayerRepository) FindByID(id uint) (*models.Player, error) {
    var player models.Player
    err := r.db.Preload("Weapons").Preload("CurrentWeapon").
        First(&player, id).Error
    return &player, err
}

func (r *PlayerRepository) FindByUsername(username string) (*models.Player, error) {
    var player models.Player
    err := r.db.Where("username = ?", username).First(&player).Error
    return &player, err
}

func (r *PlayerRepository) Update(player *models.Player) error {
    return r.db.Save(player).Error
}

func (r *PlayerRepository) UpdateGold(id uint, gold int64) error {
    return r.db.Model(&models.Player{}).
        Where("id = ?", id).
        Update("gold", gold).Error
}

func (r *PlayerRepository) FindTopPlayersByLevel(limit int) ([]models.Player, error) {
    var players []models.Player
    err := r.db.Order("level DESC, experience DESC").
        Limit(limit).
        Find(&players).Error
    return players, err
}
```

### 서비스 레이어
```go
// internal/services/player_service.go
package services

import (
    "game_eating_pizza/internal/models"
    "game_eating_pizza/internal/repository"
)

type PlayerService struct {
    playerRepo *repository.PlayerRepository
    weaponRepo *repository.WeaponRepository
}

func NewPlayerService(
    playerRepo *repository.PlayerRepository,
    weaponRepo *repository.WeaponRepository,
) *PlayerService {
    return &PlayerService{
        playerRepo: playerRepo,
        weaponRepo: weaponRepo,
    }
}

func (s *PlayerService) CreatePlayer(username string) (*models.Player, error) {
    player := &models.Player{
        Username: username,
        Level:    1,
        Gold:     0,
    }
    
    if err := s.playerRepo.Create(player); err != nil {
        return nil, err
    }
    
    // 기본 무기 생성
    defaultWeapon := &models.Weapon{
        PlayerID:    player.ID,
        Name:        "Basic Sword",
        Type:        "sword",
        AttackPower: 10,
        AttackSpeed: 1.0,
        Rarity:      "common",
        Level:       1,
    }
    
    if err := s.weaponRepo.Create(defaultWeapon); err != nil {
        return nil, err
    }
    
    player.CurrentWeaponID = &defaultWeapon.ID
    s.playerRepo.Update(player)
    
    return player, nil
}
```

---

## 마이그레이션 관리

### GORM AutoMigrate (개발용)
```go
// pkg/database/db.go
func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.Player{},
        &models.Weapon{},
        &models.Dungeon{},
        &models.Monster{},
        &models.Event{},
    )
}
```

### 프로덕션: golang-migrate 사용 권장
```bash
# 마이그레이션 파일 생성
migrate create -ext sql -dir migrations -seq create_players_table

# 마이그레이션 실행
migrate -path migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" up
```

---

## 성능 모니터링

### GORM 로깅 설정
```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info), // SQL 쿼리 로깅
})
```

### 슬로우 쿼리 감지
```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second, // 1초 이상 쿼리 로깅
            LogLevel:      logger.Warn,
        },
    ),
})
```

---

## 결론

**GORM을 강력 추천합니다.**

이유:
1. ✅ JPA와 가장 유사한 개발 경험
2. ✅ 빠른 학습 곡선
3. ✅ 풍부한 기능과 커뮤니티
4. ✅ 게임 프로젝트에 충분한 성능

성능이 매우 중요해지면 나중에 sqlc로 마이그레이션하거나, 핵심 쿼리만 sqlc로 최적화하는 하이브리드 접근도 가능합니다.

---

## 참고 자료

- [GORM 공식 문서](https://gorm.io/docs/)
- [GORM 한글 문서](https://gorm.io/ko_KR/docs/)
- [GORM GitHub](https://github.com/go-gorm/gorm)
