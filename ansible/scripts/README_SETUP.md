# Ansible 스크립트 가이드

## 📋 개요

이 디렉토리에는 Self-hosted runner 환경 설정을 위한 스크립트들이 있습니다.

## 🔧 주요 스크립트

### 1. setup-self-hosted-runner.sh
GitHub Actions Self-hosted runner를 설치합니다.

**요구사항:**
- GitHub Personal Access Token (repo, admin:repo_hook 권한)
- 올바른 Repository 형식 (owner/repository)

**실행:**
```bash
cd /home/ki/src/game/ansible/scripts
cp .env.example .env
# .env 파일 수정 (GITHUB_TOKEN, GITHUB_REPO)
./setup-self-hosted-runner.sh
```

### 2. setup-docker-permissions.sh
ki 사용자가 docker 명령어를 sudo 없이 실행할 수 있도록 권한을 설정합니다.

**필요한 이유:**
- Ansible playbook이 docker compose 명령어를 실행할 때 비밀번호 입력 없이 실행
- CI/CD 자동화를 위한 필수 설정

**실행:**
```bash
cd /home/ki/src/game/ansible/scripts
chmod +x setup-docker-permissions.sh
./setup-docker-permissions.sh
```

**설정 내용:**
- ✓ docker 그룹 생성 (이미 있으면 스킵)
- ✓ ki 사용자를 docker 그룹에 추가
- ✓ Docker 소켓 권한 설정
- ✓ 권한 검증

**적용 확인:**
```bash
# 새로운 터미널 세션에서:
docker ps

# 또는 현재 세션에서:
newgrp docker
```

## 📌 설정 순서

Self-hosted runner 환경 초기 설정 순서:

1. **Docker 설치** (이미 설치되어 있다면 스킵)
   ```bash
   # Ubuntu/Debian의 경우
   sudo apt-get update
   sudo apt-get install -y docker.io
   sudo systemctl start docker
   sudo systemctl enable docker
   ```

2. **Docker 권한 설정**
   ```bash
   ./setup-docker-permissions.sh
   ```

3. **GitHub Runner 설치**
   ```bash
   cp .env.example .env
   # .env 파일 수정
   ./setup-self-hosted-runner.sh
   ```

4. **Ansible 설치** (Self-hosted runner이 자동 설치)
   ```bash
   pip install ansible
   ```

## ✅ 배포 테스트

설정이 완료되면 다음 명령어로 테스트할 수 있습니다:

```bash
# Ansible playbook 테스트
cd /home/ki/src/game/ansible
ansible-playbook playbooks/deploy.yml -i inventory/localhost.yml

# Docker 권한 확인
docker ps

# Docker Compose 확인
cd /home/ki/src/game/server
docker compose ps
```

## 🐛 문제 해결

### "sudo: a password is required" 오류
```bash
# 해결 방법 1: setup-docker-permissions.sh 다시 실행
./setup-docker-permissions.sh
newgrp docker

# 해결 방법 2: 수동으로 docker 그룹 추가
sudo usermod -aG docker ki
newgrp docker

# 검증
docker ps
```

### "docker: command not found"
```bash
# Docker 설치 확인
which docker

# Docker 설치
sudo apt-get install -y docker.io

# Docker 데몬 시작
sudo systemctl start docker
sudo systemctl enable docker
```

### ki 사용자가 docker 그룹에 없음
```bash
# 확인
groups ki

# 추가
sudo usermod -aG docker ki

# 새로운 세션에서 적용
su - ki
```

## 📚 관련 파일

- [setup-self-hosted-runner.sh](setup-self-hosted-runner.sh) - GitHub Runner 설치 스크립트
- [setup-docker-permissions.sh](setup-docker-permissions.sh) - Docker 권한 설정 스크립트
- [.env.example](.env.example) - 환경 변수 템플릿
- [README.md](README.md) - 이 파일
