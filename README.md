# Tiny Breakers: The Beating World (타이니 브레이커즈: 심장이 뛰는 세계)

Flutter Flame과 Go를 사용한 횡스크롤 방치형 액션 게임 프로젝트입니다.

## 프로젝트 개요

**Tiny Breakers**는 사용자의 현실 활동(걷기)이 게임 내 성장으로 직결되는 방치형 게임입니다.

### 핵심 특징
- **클라이언트**: Flutter Flame 엔진 사용
- **백엔드**: Go 언어로 구현
- **플랫폼**: iOS/Android
- **특수 기능**: Apple Watch 연동, 걸음 수 연동, 멀티플레이 레이드

## 게임 플레이

- **타이니(작은 정령)**들이 거대한 무기를 들고 보석을 깨부수는 횡스롤 게임
- **걸음 수 연동**: 현실에서 걸을수록 대장간의 화로가 뜨거워져 무기 제작 속도 증가
- **자동 전투**: 타이니들이 자동으로 보석(몬스터)을 처치
- **무기 강화 시스템**: 대장간에서 무기를 제작하고 강화
- **멀티플레이 레이드**: 다른 유저들과 협력하여 거대 수정 거인(World Boss) 처치

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
