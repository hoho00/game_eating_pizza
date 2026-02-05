#!/bin/bash
# Docker 권한 설정 스크립트
# Self-hosted runner 서버에서 ki 사용자가 sudo 없이 docker를 실행할 수 있도록 설정합니다.

set -e

echo "=========================================="
echo "Docker 권한 설정 시작"
echo "=========================================="

# 1. docker 그룹 생성 (이미 있으면 스킵)
echo "1. docker 그룹 확인..."
if ! getent group docker > /dev/null; then
    echo "  docker 그룹 생성 중..."
    sudo groupadd docker
else
    echo "  ✓ docker 그룹이 이미 존재합니다"
fi

# 2. ki 사용자를 docker 그룹에 추가
echo ""
echo "2. ki 사용자를 docker 그룹에 추가 중..."
sudo usermod -aG docker ki
echo "  ✓ ki 사용자가 docker 그룹에 추가되었습니다"

# 3. Docker 데몬 소켓 권한 확인
echo ""
echo "3. Docker 소켓 권한 확인 중..."
sudo chmod 666 /var/run/docker.sock 2>/dev/null || true
echo "  ✓ Docker 소켓 권한이 설정되었습니다"

# 4. 설정 적용 (새로운 시스템 호출 필요)
echo ""
echo "4. 설정 적용..."
newgrp docker << END
echo "  ✓ 그룹 변경이 적용되었습니다"
END

# 5. 검증
echo ""
echo "5. 권한 검증 중..."
if id -nG ki | grep -qw docker; then
    echo "  ✓ ki 사용자가 docker 그룹의 멤버입니다"
else
    echo "  ✗ ki 사용자가 docker 그룹의 멤버가 아닙니다"
    echo "  새로운 쉘 세션에서 적용되려면 로그아웃 후 다시 로그인하세요"
fi

# 6. docker 명령어 테스트
echo ""
echo "6. Docker 명령어 테스트 중..."
if docker ps > /dev/null 2>&1; then
    echo "  ✓ Docker 명령어가 정상 작동합니다"
else
    echo "  ✗ Docker 명령어 테스트 실패"
    echo "  다음 명령어를 실행하세요:"
    echo "    su - ki"
fi

echo ""
echo "=========================================="
echo "✅ Docker 권한 설정이 완료되었습니다!"
echo "=========================================="
echo ""
echo "주의: 새로운 터미널 세션에서 변경사항이 적용됩니다."
echo "현재 세션에서는 다음 명령어로 적용할 수 있습니다:"
echo "  newgrp docker"
