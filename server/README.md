# Game Eating Pizza - Go Backend Server

횡스크롤 방치형 게임의 백엔드 서버입니다.

## 기술 스택

- **언어**: Go 1.21+
- **웹 프레임워크**: Gin
- **ORM**: GORM
- **데이터베이스**: PostgreSQL / MySQL
- **인증**: JWT (구현 예정)
- **API 문서**: Swagger (swaggo)

## 프로젝트 구조

```
server/
├── cmd/
│   └── server/
│       └── main.go          # 서버 진입점
├── internal/
│   ├── api/                 # API 핸들러 및 라우팅
│   │   ├── handlers/        # HTTP 핸들러
│   │   └── middleware/      # 미들웨어
│   ├── config/              # 설정 관리
│   ├── models/              # 데이터 모델
│   ├── repository/          # 데이터 접근 계층
│   └── services/            # 비즈니스 로직
└── pkg/
    └── database/            # 데이터베이스 연결
```

## 시작하기

### 1. 사전 요구사항

- Go 1.21 이상
- PostgreSQL 또는 MySQL (또는 Mock DB 사용)
- Git
- Swagger CLI (선택사항, API 문서 생성용)

### 2. 환경 설정

```bash
# .env.example을 .env로 복사
cp .env.example .env

# .env 파일을 편집하여 데이터베이스 설정 수정
```

### 3. 의존성 설치

```bash
go mod download
```

### 4. 데이터베이스 설정

PostgreSQL 예시:
```bash
createdb game_db
```

또는 MySQL:
```sql
CREATE DATABASE game_db;
```

### 5. 서버 실행

```bash
# 개발 모드
go run cmd/server/main.go

# 또는 빌드 후 실행
go build -o bin/server cmd/server/main.go
./bin/server
```

서버가 `http://localhost:8080`에서 실행됩니다.

## Swagger API 문서

### Swagger 문서 생성

Swagger 문서를 생성하려면:

1. **swag CLI 설치**:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

**PATH 설정** (필요한 경우):
```bash
# ~/.zshrc 또는 ~/.bashrc에 추가
export PATH=$PATH:$(go env GOPATH)/bin
# 또는
export PATH=$PATH:$HOME/go/bin

# 적용
source ~/.zshrc
```

2. **문서 생성**:
```bash
cd server
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

또는 Makefile 사용 (자동 설치 포함):
```bash
make swagger
```

**문제 해결**: `command not found: swag` 에러가 발생하면 `SWAGGER_SETUP.md` 파일을 참고하세요.

3. **서버 실행 후 접속**:
```
http://localhost:8080/swagger/index.html
```

### Swagger UI 사용

- Swagger UI는 서버 실행 시 자동으로 `/swagger/*any` 경로에서 제공됩니다
- 브라우저에서 `http://localhost:8080/swagger/index.html` 접속
- API 엔드포인트를 직접 테스트할 수 있습니다
- 인증이 필요한 API는 우측 상단의 "Authorize" 버튼을 클릭하여 토큰을 입력하세요

## API 엔드포인트

### Health Check
- `GET /health` - 서버 상태 확인

### 인증
- `POST /api/v1/auth/register` - 회원가입
- `POST /api/v1/auth/login` - 로그인
- `POST /api/v1/auth/refresh` - 토큰 갱신

### 플레이어 (인증 필요)
- `GET /api/v1/players/me` - 내 정보 조회
- `PUT /api/v1/players/me` - 내 정보 수정
- `GET /api/v1/players/leaderboard` - 리더보드

### 무기 (인증 필요)
- `GET /api/v1/weapons` - 무기 목록
- `POST /api/v1/weapons` - 무기 생성
- `PUT /api/v1/weapons/:id/upgrade` - 무기 강화
- `PUT /api/v1/weapons/:id/equip` - 무기 장착

### 던전 (인증 필요)
- `GET /api/v1/dungeons` - 던전 목록
- `GET /api/v1/dungeons/:id` - 던전 상세
- `POST /api/v1/dungeons/:id/enter` - 던전 입장
- `POST /api/v1/dungeons/:id/clear` - 던전 클리어

## 개발

### 데이터베이스 마이그레이션

개발 환경에서는 GORM의 AutoMigrate를 사용할 수 있습니다:

```go
// main.go에서
database.AutoMigrate(db, &models.Player{}, &models.Weapon{}, &models.Dungeon{})
```

프로덕션 환경에서는 `golang-migrate` 같은 마이그레이션 도구 사용을 권장합니다.

### 테스트

```bash
go test ./...
```

## TODO

- [ ] JWT 인증 구현
- [ ] 비밀번호 해시 검증
- [x] API 문서화 (Swagger)
- [ ] 단위 테스트 작성
- [ ] 통합 테스트 작성
- [ ] 로깅 시스템 개선
- [ ] 에러 핸들링 개선
- [ ] 데이터베이스 마이그레이션 도구 설정

## 라이선스

[라이선스 정보를 여기에 추가하세요]
