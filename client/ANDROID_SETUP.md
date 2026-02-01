# Android 가상 머신에서 실행하기

## 1. Android Studio에서 에뮬레이터 생성

### 방법 1: Android Studio AVD Manager 사용

1. **Android Studio 실행**
2. **Tools → Device Manager** (또는 **More Actions → Virtual Device Manager**)
3. **Create Device** 클릭
4. **기기 선택**: 
   - 권장: **Pixel 5** 또는 **Pixel 6**
   - 또는 원하는 기기 선택
5. **시스템 이미지 선택**:
   - **API Level 33 (Android 13)** 또는 **API Level 34 (Android 14)** 권장
   - **x86_64** 아키텍처 선택 (Intel Mac의 경우)
   - **arm64-v8a** 아키텍처 선택 (Apple Silicon Mac의 경우)
6. **Finish** 클릭

### 방법 2: 명령어로 확인

```bash
# 사용 가능한 에뮬레이터 목록 확인
flutter emulators

# 특정 에뮬레이터 실행
flutter emulators --launch <emulator_id>
```

## 2. 에뮬레이터 실행

### Android Studio에서
1. **Device Manager**에서 생성한 에뮬레이터 옆의 **▶️ Play** 버튼 클릭
2. 에뮬레이터가 부팅될 때까지 대기 (1-2분)

### 명령어로
```bash
# 에뮬레이터 목록 확인
flutter emulators

# 에뮬레이터 실행 (예: Pixel_5_API_33)
flutter emulators --launch Pixel_5_API_33
```

## 3. Flutter 프로젝트 실행

### 의존성 설치
```bash
cd client
flutter pub get
```

### Android 에뮬레이터에서 실행
```bash
# 에뮬레이터가 실행 중인지 확인
flutter devices

# 실행 (에뮬레이터가 하나만 있으면 자동 선택)
flutter run

# 또는 특정 에뮬레이터 지정
flutter run -d <device-id>
```

## 4. 문제 해결

### 에뮬레이터가 보이지 않는 경우
```bash
# Flutter doctor로 문제 확인
flutter doctor

# Android 라이선스 동의 (필요한 경우)
flutter doctor --android-licenses
```

### ADB 연결 문제
```bash
# ADB 재시작
adb kill-server
adb start-server

# 연결된 기기 확인
adb devices
```

### 빌드 에러
```bash
# 클린 빌드
flutter clean
flutter pub get
flutter run
```

## 5. 빠른 시작 명령어

```bash
# 1. client 폴더로 이동
cd client

# 2. 의존성 설치
flutter pub get

# 3. 에뮬레이터 확인
flutter devices

# 4. 실행
flutter run
```

## 6. Hot Reload 사용

앱이 실행 중일 때:
- **r**: Hot Reload (빠른 재시작)
- **R**: Hot Restart (완전 재시작)
- **q**: 종료

## 참고

- 첫 실행 시 빌드 시간이 오래 걸릴 수 있습니다 (1-3분)
- 에뮬레이터는 최소 2GB RAM을 권장합니다
- Apple Silicon Mac의 경우 arm64 에뮬레이터를 사용하세요
