# 프로젝트 구조: Tiny Breakers

## 전체 프로젝트 구조

```
tiny_breakers/ (또는 game_eating_pizza/)
├── client/                    # Flutter Flame 클라이언트
│   ├── lib/
│   │   ├── main.dart
│   │   ├── game/
│   │   │   ├── game.dart          # 메인 게임 클래스
│   │   │   ├── components/         # 게임 컴포넌트
│   │   │   │   ├── player.dart
│   │   │   │   ├── monster.dart
│   │   │   │   ├── weapon.dart
│   │   │   │   └── projectile.dart
│   │   │   ├── systems/           # 게임 시스템
│   │   │   │   ├── combat_system.dart
│   │   │   │   ├── spawn_system.dart
│   │   │   │   └── upgrade_system.dart
│   │   │   └── managers/           # 게임 매니저
│   │   │       ├── game_manager.dart
│   │   │       └── resource_manager.dart
│   │   ├── models/                # 데이터 모델
│   │   │   ├── player_model.dart
│   │   │   ├── weapon_model.dart
│   │   │   ├── monster_model.dart
│   │   │   └── dungeon_model.dart
│   │   ├── services/              # 서비스 레이어
│   │   │   ├── api_service.dart
│   │   │   ├── auth_service.dart
│   │   │   └── storage_service.dart
│   │   ├── screens/               # UI 화면
│   │   │   ├── home_screen.dart
│   │   │   ├── game_screen.dart
│   │   │   ├── upgrade_screen.dart
│   │   │   └── dungeon_screen.dart
│   │   ├── widgets/               # 재사용 가능한 위젯
│   │   │   ├── health_bar.dart
│   │   │   ├── gold_display.dart
│   │   │   └── weapon_card.dart
│   │   └── utils/                 # 유틸리티
│   │       ├── constants.dart
│   │       └── helpers.dart
│   ├── watch_extension/           # Apple Watch Extension
│   │   └── lib/
│   │       ├── main.dart
│   │       └── watch_screens/
│   ├── assets/                    # 게임 리소스
│   │   ├── images/
│   │   ├── audio/
│   │   └── data/
│   ├── pubspec.yaml
│   └── README.md
│
├── server/                        # Go 백엔드 서버
│   ├── cmd/
│   │   └── server/
│   │       └── main.go            # 서버 진입점
│   ├── internal/
│   │   ├── api/                   # API 핸들러
│   │   │   ├── handlers/
│   │   │   │   ├── auth_handler.go
│   │   │   │   ├── player_handler.go
│   │   │   │   ├── weapon_handler.go
│   │   │   │   ├── dungeon_handler.go
│   │   │   │   └── event_handler.go
│   │   │   ├── middleware/
│   │   │   │   ├── auth_middleware.go
│   │   │   │   └── cors_middleware.go
│   │   │   └── routes.go
│   │   ├── models/                # 데이터 모델
│   │   │   ├── player.go
│   │   │   ├── weapon.go
│   │   │   ├── monster.go
│   │   │   └── dungeon.go
│   │   ├── services/              # 비즈니스 로직
│   │   │   ├── auth_service.go
│   │   │   ├── player_service.go
│   │   │   ├── weapon_service.go
│   │   │   ├── dungeon_service.go
│   │   │   └── event_service.go
│   │   ├── repository/            # 데이터 접근 계층
│   │   │   ├── player_repository.go
│   │   │   ├── weapon_repository.go
│   │   │   └── dungeon_repository.go
│   │   └── config/                # 설정
│   │       └── config.go
│   ├── pkg/                       # 공용 패키지
│   │   ├── database/
│   │   │   └── db.go
│   │   ├── jwt/
│   │   │   └── jwt.go
│   │   └── logger/
│   │       └── logger.go
│   ├── migrations/                # DB 마이그레이션
│   ├── docker-compose.yml         # Docker Compose로 관리되는 PostgreSQL 설정 (데이터 지속성 포함)
│   ├── Dockerfile
│   ├── .env.example               # 환경 변수 예시
│   ├── go.mod
│   ├── go.sum
│   └── README.md
│
├── docs/                          # 문서
│   ├── GAME_DESIGN_DOCUMENT.md
│   ├── API_DOCUMENTATION.md
│   └── DEPLOYMENT.md
│
├── scripts/                       # 유틸리티 스크립트
│   ├── setup.sh
│   └── deploy.sh
│
└── README.md                      # 프로젝트 루트 README
```

## 주요 디렉토리 설명

### Client (Flutter Flame)
- **game/**: 게임 엔진 관련 코드
- **models/**: 데이터 모델 정의
- **services/**: API 통신 및 로컬 저장소 관리
- **screens/**: UI 화면
- **watch_extension/**: Apple Watch 전용 코드

### Server (Go)
- **cmd/server/**: 애플리케이션 진입점
- **internal/**: 내부 패키지 (외부에서 import 불가)
- **pkg/**: 공용 패키지 (외부에서 import 가능)
- **migrations/**: 데이터베이스 스키마 변경 이력

## 개발 환경 설정

### 필수 도구
- Flutter SDK (최신 안정 버전)
- Go 1.21 이상
- PostgreSQL 또는 MySQL
- Docker (선택사항)

### 권장 IDE
- VS Code 또는 Android Studio (Flutter)
- GoLand 또는 VS Code (Go)
