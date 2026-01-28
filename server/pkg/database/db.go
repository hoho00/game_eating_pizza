package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"game_eating_pizza/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB는 전역 데이터베이스 연결 인스턴스
var DB *gorm.DB

// Connect는 데이터베이스에 연결하고 전역 DB 인스턴스를 초기화합니다
func Connect(cfg *config.Config) (*gorm.DB, error) {
	var dsn string
	var dialector gorm.Dialector

	switch cfg.DBDriver {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
		)
		dialector = postgres.Open(dsn)
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		)
		dialector = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.DBDriver)
	}

	// GORM 로거 설정
	logLevel := logger.Info
	if cfg.Env == "production" {
		logLevel = logger.Error
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 1초 이상 쿼리는 슬로우 쿼리로 표시
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  cfg.Env != "production",
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 연결 풀 설정
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// 연결 풀 최대 연결 수
	sqlDB.SetMaxOpenConns(25)
	// 연결 풀 최대 유휴 연결 수
	sqlDB.SetMaxIdleConns(5)
	// 연결 최대 수명
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = db
	return db, nil
}

// AutoMigrate는 모든 모델을 자동으로 마이그레이션합니다
// 개발 환경에서만 사용하고, 프로덕션에서는 golang-migrate 사용 권장
func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}

// Close는 데이터베이스 연결을 닫습니다
func Close() error {
	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// HealthCheck는 데이터베이스 연결 상태를 확인합니다
func HealthCheck() error {
	if DB == nil {
		return fmt.Errorf("database connection is nil")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}
