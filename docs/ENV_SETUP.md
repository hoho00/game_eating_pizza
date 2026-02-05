# .env 파일 설정 가이드

## 📝 개요

이 프로젝트는 `.env` 파일을 사용하여 환경 변수를 관리합니다.

## 🔧 설정 방법

### 1. Self-hosted Runner 환경

**GitHub Actions self-hosted runner에서 자동으로 .env 파일을 적용합니다:**

```bash
# 서버의 게임 소스 디렉토리에 .env 파일 생성
cd /home/ki/src/game
cat > .env << 'EOF'
# 필요한 환경 변수들 입력
DATABASE_URL=...
API_KEY=...
EOF
```

**CI/CD 배포 프로세스:**
1. GitHub Actions가 dev 브랜치 push 감지
2. 바이너리 빌드 후 `.env` 파일을 `server/` 디렉토리로 자동 복사
3. Ansible이 서버에 파일 배포

### 2. Local 개발 환경

```bash
# 프로젝트 루트에 .env 파일 생성
cat > .env << 'EOF'
# 로컬 개발용 환경 변수
DATABASE_URL=local_db_url
API_KEY=local_api_key
EOF
```

### 3. .env 파일 위치

```
/home/ki/src/game/
├── .env                    ← 여기에 생성 (Git에 미포함)
├── server/
│   ├── server             ← 빌드된 바이너리
│   ├── server-batch       ← 배치 바이너리
│   └── .env               ← CI/CD에서 자동 복사
├── ansible/
│   └── playbooks/
│       ├── deploy.yml
│       └── deploy-localhost.yml
└── .gitignore             ← .env 제외됨
```

## ✅ 자동 배포 흐름

```
Push to dev branch
    ↓
GitHub Actions (self-hosted runner)
    ↓
1. 소스 체크아웃
2. Go 바이너리 빌드
3. .env 파일 server/ 디렉토리로 복사 ⭐
4. Ansible 배포 실행
    ↓
Ansible playbook
    ↓
1. 서버 디렉토리 생성
2. .env 파일 존재 여부 확인 ✓
3. .env 파일 복사 (존재하면)
4. systemd 서비스 생성 및 재시작
    ↓
배포 완료 ✅
```

## 📌 주의사항

- **절대 .env 파일을 Git에 커밋하지 마세요** (.gitignore에 이미 포함됨)
- **민감한 정보** (API 키, 데이터베이스 URL 등)는 반드시 .env 파일에만 저장
- Self-hosted runner 서버에 .env 파일이 있어야 배포 시 자동으로 적용됨

## 🔍 배포 상태 확인

```bash
# 서버에서 .env 파일 확인
cat ~/backend/game-server/.env

# 서비스 상태 확인
sudo systemctl status tiny-breakers-game

# 로그 확인
sudo journalctl -u tiny-breakers-game -f
```
