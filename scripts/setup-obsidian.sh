#!/bin/bash
# Obsidian 환경 초기 설정 스크립트
# 개발 서버에서 직접 실행하거나 Ansible로 실행

set -e

# 스크립트가 있는 디렉토리로 이동
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

echo "Obsidian 환경 설정을 시작합니다..."
echo "프로젝트 루트: $PROJECT_ROOT"

# 프로젝트 루트로 이동
cd "$PROJECT_ROOT"

# Ansible playbook 경로 확인
PLAYBOOK_PATH="$PROJECT_ROOT/ansible/playbooks/setup-obsidian.yml"
INVENTORY_PATH="$PROJECT_ROOT/ansible/inventory/localhost.yml"

if [ ! -f "$PLAYBOOK_PATH" ]; then
    echo "오류: Ansible playbook을 찾을 수 없습니다: $PLAYBOOK_PATH"
    echo "현재 디렉토리: $(pwd)"
    echo "파일 목록:"
    ls -la ansible/playbooks/ 2>/dev/null || echo "ansible/playbooks/ 디렉토리가 없습니다"
    exit 1
fi

if [ ! -f "$INVENTORY_PATH" ]; then
    echo "오류: Ansible inventory를 찾을 수 없습니다: $INVENTORY_PATH"
    exit 1
fi

# Ansible 설치 확인
if ! command -v ansible-playbook &> /dev/null; then
    echo "Ansible이 설치되어 있지 않습니다. 설치를 시작합니다..."
    
    # Python3 확인
    if ! command -v python3 &> /dev/null; then
        echo "오류: Python3가 설치되어 있지 않습니다."
        echo "다음 명령으로 설치하세요:"
        echo "  Ubuntu/Debian: sudo apt update && sudo apt install -y python3 python3-pip"
        echo "  CentOS/RHEL: sudo yum install -y python3 python3-pip"
        exit 1
    fi
    
    # pip3로 Ansible 설치
    echo "pip3를 사용하여 Ansible을 설치합니다..."
    python3 -m pip install --user ansible || {
        echo "pip3 설치 실패. sudo 권한으로 시도합니다..."
        sudo python3 -m pip install ansible || {
            echo "오류: Ansible 설치에 실패했습니다."
            echo "수동으로 설치하세요:"
            echo "  pip3 install --user ansible"
            echo "  또는"
            echo "  sudo pip3 install ansible"
            exit 1
        }
    }
    
    # PATH에 ~/.local/bin 추가 (필요한 경우)
    if [ -d "$HOME/.local/bin" ] && [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        export PATH="$HOME/.local/bin:$PATH"
        echo "PATH에 ~/.local/bin을 추가했습니다."
    fi
    
    # 다시 확인
    if ! command -v ansible-playbook &> /dev/null; then
        echo "오류: Ansible 설치 후에도 명령을 찾을 수 없습니다."
        echo "다음 명령을 실행하세요:"
        echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
        echo "  또는 터미널을 재시작하세요."
        exit 1
    fi
    
    echo "Ansible 설치가 완료되었습니다."
fi

echo "Ansible을 사용하여 설정합니다..."
cd "$PROJECT_ROOT/ansible"
ansible-playbook playbooks/setup-obsidian.yml -i inventory/localhost.yml

echo ""
echo "설정이 완료되었습니다!"
echo "문서는 http://100.69.72.59 에서 확인할 수 있습니다."
