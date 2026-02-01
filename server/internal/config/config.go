package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config는 애플리케이션 설정을 담는 구조체입니다
type Config struct {
	// 서버 설정
	ServerPort string
	ServerHost string
	Env        string // "development", "production"

	// 데이터베이스 설정
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBDriver   string // "postgres" or "mysql"
	DBSSLMode  string

	// JWT 설정
	JWTSecret     string
	JWTExpiration int // 시간 (시간 단위)

	// CORS 설정
	CORSAllowedOrigins []string

	// Redis 설정 (캐싱, 세션, 실시간 데이터용)
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int // Redis 데이터베이스 번호 (0-15)
}

var AppConfig *Config

// LoadConfig는 환경 변수에서 설정을 로드합니다
func LoadConfig() (*Config, error) {
	// .env 파일 로드 (파일이 없어도 에러 무시)
	_ = godotenv.Load()

	config := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerHost: getEnv("SERVER_HOST", "0.0.0.0"),
		Env:        getEnv("ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "game_db"),
		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 24),

		CORSAllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""), // 비밀번호가 설정되어 있어야 합니다
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
	}

	AppConfig = config
	return config, nil
}

// getEnv는 환경 변수를 가져오고, 없으면 기본값을 반환합니다
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt는 환경 변수를 정수로 변환합니다
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getEnvAsSlice는 환경 변수를 슬라이스로 변환합니다 (쉼표로 구분)
func getEnvAsSlice(key string, defaultValue []string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	// 쉼표로 구분된 문자열을 슬라이스로 변환
	result := []string{}
	start := 0
	for i, char := range valueStr {
		if char == ',' {
			if i > start {
				result = append(result, valueStr[start:i])
			}
			start = i + 1
		}
	}
	if start < len(valueStr) {
		result = append(result, valueStr[start:])
	}
	if len(result) == 0 {
		return defaultValue
	}
	return result
}

// getEnvAsBool는 환경 변수를 불린으로 변환합니다
func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	return valueStr == "true" || valueStr == "1" || valueStr == "yes"
}
