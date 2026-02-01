#!/bin/bash

# Android 에뮬레이터에서 Flutter 앱 실행 스크립트

echo "🚀 Android 에뮬레이터에서 Flutter 앱 실행하기"
echo ""

# 1. 의존성 설치
echo "📦 의존성 설치 중..."
flutter pub get

# 2. 에뮬레이터 확인
echo ""
echo "📱 사용 가능한 에뮬레이터 확인 중..."
flutter emulators

# 3. 에뮬레이터 실행 (이미 실행 중이면 스킵)
echo ""
echo "📱 에뮬레이터 실행 중..."
flutter emulators --launch Medium_Phone_API_VanillaIceCream

# 4. 에뮬레이터가 부팅될 때까지 대기
echo ""
echo "⏳ 에뮬레이터 부팅 대기 중... (30초)"
sleep 30

# 5. 연결된 기기 확인
echo ""
echo "🔍 연결된 기기 확인 중..."
flutter devices

# 6. 앱 실행
echo ""
echo "🎮 Flutter 앱 실행 중..."
flutter run -d android
