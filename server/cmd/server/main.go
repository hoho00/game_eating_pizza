package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"game_eating_pizza/internal/config"
	"game_eating_pizza/internal/api"
	"game_eating_pizza/pkg/database"
	"gorm.io/gorm"

	_ "game_eating_pizza/docs" // Swagger 문서를 위한 import
)

// @title           Game Eating Pizza API
// @version         1.0
// @description     횡스크롤 방치형 게임 백엔드 API 서버
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// 설정 로드
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 데이터베이스 연결 (Mock 사용 시에는 연결하지 않음)
	var db *gorm.DB
	if !cfg.UseMockDB {
		var err error
		db, err = database.Connect(cfg)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		defer database.Close()
		log.Println("Database connected successfully")
	} else {
		log.Println("Using Mock Database (MVP development mode)")
	}

	// 라우터 설정
	router := api.SetupRouter(db, cfg)

	// 서버 시작
	serverAddr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Graceful shutdown을 위한 채널
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 서버를 고루틴에서 시작
	go func() {
		log.Printf("Server starting on %s", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 종료 신호 대기
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown (5초 타임아웃)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
