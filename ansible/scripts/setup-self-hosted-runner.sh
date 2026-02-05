#!/bin/bash
# GitHub Actions Self-hosted Runner 설치 스크립트
# 개발 서버(100.69.72.59)에서 실행

set -e

# .env 파일이 있으면 로드
if [ -f "$(dirname "$0")/.env" ]; then
    source "$(dirname "$0")/.env"
fi

echo "GitHub Actions Self-hosted Runner 설치를 시작합니다..."

# 환경변수에서 값 가져오기
GITHUB_TOKEN="${GITHUB_TOKEN:-}"
GITHUB_REPO="${GITHUB_REPO:-}"

# 필수 환경변수 확인
if [ -z "$GITHUB_TOKEN" ]; then
    echo "에러: GITHUB_TOKEN 환경변수가 설정되지 않았습니다."
    echo "설정 방법: .env 파일을 생성하거나 환경변수를 설정하세요."
    exit 1
fi

if [ -z "$GITHUB_REPO" ]; then
    echo "에러: GITHUB_REPO 환경변수가 설정되지 않았습니다."
    echo "설정 방법: .env 파일을 생성하거나 환경변수를 설정하세요."
    exit 1
fi

# Runner 버전
RUNNER_VERSION="2.311.0"
ARCH="x64"

# Runner 설치 디렉토리
RUNNER_DIR="$HOME/actions-runner"
SERVICE_NAME="github-actions-runner"

echo "설정 정보:"
echo "  Repository: $GITHUB_REPO"
echo "  Runner Directory: $RUNNER_DIR"

# 기존 Runner 제거 (있는 경우)
if [ -d "$RUNNER_DIR" ]; then
    echo "기존 Runner를 제거합니다..."
    sudo systemctl stop $SERVICE_NAME || true
    sudo systemctl disable $SERVICE_NAME || true
    cd $RUNNER_DIR
    sudo ./svc.sh uninstall || true
    cd ~
    rm -rf $RUNNER_DIR
fi

# Runner 다운로드
echo "Runner를 다운로드합니다..."
mkdir -p $RUNNER_DIR
cd $RUNNER_DIR

curl -o actions-runner-linux-${ARCH}-${RUNNER_VERSION}.tar.gz -L https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-${ARCH}-${RUNNER_VERSION}.tar.gz
tar xzf ./actions-runner-linux-${ARCH}-${RUNNER_VERSION}.tar.gz

# Runner 구성
echo "Runner를 구성합니다..."
./config.sh --url https://github.com/${GITHUB_REPO} --token ${GITHUB_TOKEN} --name dev-server-runner --work _work --unattended

# Service로 설치
echo "Runner를 서비스로 설치합니다..."
sudo ./svc.sh install ki
sudo ./svc.sh start
sudo ./svc.sh status

echo "설치가 완료되었습니다!"
echo "Runner 상태 확인: sudo systemctl status $SERVICE_NAME"
echo "Runner 로그 확인: sudo journalctl -u $SERVICE_NAME -f"
