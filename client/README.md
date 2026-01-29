# Game Eating Pizza - Flutter Client

횡스크롤 방치형 게임의 Flutter Flame 클라이언트입니다.

## 기술 스택

- **프레임워크**: Flutter 3.0+
- **게임 엔진**: Flame 1.15+
- **상태 관리**: Provider
- **네트워크**: Dio
- **로컬 저장소**: SharedPreferences

## 프로젝트 구조

```
lib/
├── main.dart              # 앱 진입점
├── game/                  # 게임 엔진 관련
│   ├── game.dart          # 메인 게임 클래스
│   ├── components/        # 게임 컴포넌트
│   ├── systems/           # 게임 시스템
│   └── managers/          # 게임 매니저
├── models/                # 데이터 모델
├── services/              # 서비스 레이어
├── screens/               # UI 화면
├── widgets/               # 재사용 가능한 위젯
└── utils/                 # 유틸리티
```

## 시작하기

### 1. 사전 요구사항

- Flutter SDK 3.0 이상
- Dart 3.0 이상
- iOS Simulator 또는 Android Emulator

### 2. 의존성 설치

```bash
cd client
flutter pub get
```

### 3. 실행

```bash
# iOS
flutter run -d ios

# Android
flutter run -d android

# 특정 디바이스
flutter devices
flutter run -d <device-id>
```

## 개발 가이드

### 게임 구조

- **GameEatingPizza**: 메인 게임 클래스 (FlameGame 상속)
- **GameManager**: 게임 상태 관리
- **Player**: 플레이어 컴포넌트
- **Monster**: 몬스터 컴포넌트
- **SpawnSystem**: 몬스터 스폰 시스템
- **CombatSystem**: 전투 시스템

### 서버 연동

서버 기본 URL: `http://localhost:8080/api/v1`

개발 시 실제 서버 주소로 변경:
```dart
// lib/services/api_service.dart
ApiService(baseUrl: 'http://your-server-ip:8080/api/v1')
```

## 다음 단계

- [ ] 플레이어 스프라이트 추가
- [ ] 몬스터 스프라이트 추가
- [ ] 배경 스크롤 구현
- [ ] 충돌 감지 시스템
- [ ] 공격 애니메이션
- [ ] UI 개선

## 참고 자료

- [Flutter 공식 문서](https://flutter.dev/docs)
- [Flame 공식 문서](https://docs.flame-engine.org/)
- [Provider 패키지](https://pub.dev/packages/provider)
