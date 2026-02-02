package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"game_eating_pizza/internal/config"
	"game_eating_pizza/pkg/database"
)

// 배치 스케줄링 서버 진입점
// 주기적으로 실행되는 배치 작업을 처리합니다
func main() {
	// 설정 로드
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 데이터베이스 연결
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Println("Database connected successfully")
	
	// db 변수는 배치 작업 구현 시 사용됩니다
	_ = db

	// 배치 작업 실행 컨텍스트
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown을 위한 채널
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 배치 작업 실행 (예시)
	go func() {
		ticker := time.NewTicker(1 * time.Minute) // 1분마다 실행
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// TODO: 배치 작업 구현
				// 예: 던전 상태 업데이트, 이벤트 생성, 통계 수집 등
				log.Println("Batch job executed")
			}
		}
	}()

	// 종료 신호 대기
	<-quit
	log.Println("Shutting down batch server...")

	// Graceful shutdown
	cancel()
	time.Sleep(1 * time.Second)

	log.Println("Batch server exited")
}
