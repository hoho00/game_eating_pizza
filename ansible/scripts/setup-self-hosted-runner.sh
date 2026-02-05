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
    echo "         GitHub Personal Access Token이 필요합니다."
    echo "         - repo 권한 포함"
    echo "         - admin:repo_hook 권한 포함"
    exit 1
fi

if [ -z "$GITHUB_REPO" ]; then
    echo "에러: GITHUB_REPO 환경변수가 설정되지 않았습니다."
    echo "설정 방법: .env 파일을 생성하거나 환경변수를 설정하세요."
    echo "형식: owner/repository (예: username/repo-name)"
    exit 1
fi

# GITHUB_REPO 형식 검증 (owner/repo 형식이어야 함)
if [[ ! "$GITHUB_REPO" =~ ^[a-zA-Z0-9_-]+/[a-zA-Z0-9_.-]+$ ]]; then
    echo "에러: GITHUB_REPO 형식이 잘못되었습니다."
    echo "올바른 형식: owner/repository"
    echo "예시: myusername/my-repo"
    exit 1
fi

# Token 유효성 검증
echo "GitHub Token 유효성을 검증하고 있습니다..."
TOKEN_VALID=$(curl -s -H "Authorization: token ${GITHUB_TOKEN}" https://api.github.com/user | grep -c "login" || true)
if [ "$TOKEN_VALID" -eq 0 ]; then
    echo "에러: GitHub Token이 유효하지 않습니다."
    echo "다음을 확인하세요:"
    echo "  1. Token이 만료되지 않았는지 확인"
    echo "  2. Token에 repo, admin:repo_hook 권한이 있는지 확인"
    echo "  3. Token 값이 올바른지 확인"
    exit 1
fi
echo "✓ Token이 유효합니다."

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
echo "Repository: https://github.com/${GITHUB_REPO}"

# config.sh 실행 및 오류 처리
if ./config.sh --url https://github.com/${GITHUB_REPO} --token ${GITHUB_TOKEN} --name dev-server-runner --work _work --unattended; then
    echo "✓ Runner 구성이 완료되었습니다."
else
    echo "에러: Runner 구성에 실패했습니다."
    echo "다음을 확인하세요:"
    echo "  1. GITHUB_TOKEN이 유효한지 확인 (repo, admin:repo_hook 권한 필요)"
    echo "  2. GITHUB_REPO 형식이 올바른지 확인 (owner/repository)"
    echo "  3. Repository가 실제로 존재하는지 확인"
    echo "  4. Token에 해당 Repository 접근 권한이 있는지 확인"
    exit 1
fi

# Service로 설치
echo "Runner를 서비스로 설치합니다..."
if sudo ./svc.sh install ki; then
    echo "✓ 서비스 설치 완료"
else
    echo "에러: 서비스 설치 실패"
    exit 1
fi

if sudo ./svc.sh start; then
    echo "✓ 서비스 시작 완료"
else
    echo "에러: 서비스 시작 실패"
    exit 1
fi

sudo ./svc.sh status
echo ""
echo "✅ 설치가 완료되었습니다!"
echo ""
echo "다음 명령어로 Runner 상태를 확인할 수 있습니다:"
echo "  상태 확인: sudo systemctl status $SERVICE_NAME"
echo "  로그 확인: sudo journalctl -u $SERVICE_NAME -f"
echo ""
echo "GitHub Actions 확인:"
echo "  https://github.com/${GITHUB_REPO}/settings/actions/runners"
