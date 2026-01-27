# 횡스크롤 방치형 액션 게임

Flutter Flame과 Go를 사용한 횡스크롤 방치형 액션 게임 프로젝트입니다.

## 프로젝트 개요

이 프로젝트는 다음과 같은 특징을 가진 모바일 게임입니다:
- **클라이언트**: Flutter Flame 엔진 사용
- **백엔드**: Go 언어로 구현
- **플랫폼**: iOS/Android
- **특수 기능**: Apple Watch 연동

## 게임 플레이

- 캐릭터가 자동으로 횡스크롤하며 진행
- 몬스터를 자동/수동으로 처치
- 무기 및 캐릭터 강화 시스템
- 서버에서 제공하는 이벤트 및 던전 시스템

## 프로젝트 구조

```
game_eating_pizza/
├── client/          # Flutter Flame 클라이언트
├── server/          # Go 백엔드 서버
├── docs/            # 문서
└── scripts/         # 유틸리티 스크립트
```

## 시작하기

### 사전 요구사항

- Flutter SDK (최신 안정 버전)
- Go 1.21 이상
- PostgreSQL 또는 MySQL
- Docker (선택사항)

### 클라이언트 실행

```bash
cd client
flutter pub get
flutter run
```

### 서버 실행

```bash
cd server
go mod download
go run cmd/server/main.go
```

## 문서

- [게임 기획서](./docs/GAME_DESIGN_DOCUMENT.md)
- [향후 기획 방향성](./docs/FUTURE_PLANNING.md)
- [프로젝트 구조](./PROJECT_STRUCTURE.md)
- [Go ORM 추천 가이드](./docs/ORM_RECOMMENDATION.md)

## 개발 로드맵

자세한 개발 계획은 [게임 기획서](./docs/GAME_DESIGN_DOCUMENT.md)의 "개발 단계별 로드맵" 섹션을 참고하세요.

## 라이선스

[라이선스 정보를 여기에 추가하세요]
