# GitHub Actions Self-hosted Runner 설정

이 스크립트는 개발 서버에 GitHub Actions Self-hosted Runner를 설치합니다.

## 📋 사전 요구사항

- Ubuntu/Linux 서버
- sudo 권한
- Git, curl이 설치되어 있어야 함

## 🔧 설정 단계

### 1. GitHub Personal Access Token 생성

1. GitHub에 로그인
2. Settings → Developer settings → Personal access tokens → Tokens (classic)
3. "Generate new token" 클릭
4. 다음 권한 선택:
   - ✅ `repo` (전체)
   - ✅ `admin:repo_hook`
5. Token 생성 및 복사

### 2. .env 파일 생성

```bash
cd /home/ki/src/game/ansible/scripts
cp .env.example .env
```

`.env` 파일을 열고 다음 값을 입력:

```dotenv
GITHUB_TOKEN=your_personal_access_token_here
GITHUB_REPO=your_username/your_repo_name
```

**예시:**
```dotenv
GITHUB_TOKEN=ghp_abcdefghijklmnopqrstuvwxyz1234567890
GITHUB_REPO=myusername/my-game-repo
```

### 3. 스크립트 실행

```bash
chmod +x setup-self-hosted-runner.sh
./setup-self-hosted-runner.sh
```

## ⚠️ 일반적인 오류 해결

### HTTP 404 Not Found 오류

**원인:**
- GitHub Token이 유효하지 않음
- Token이 만료됨
- Token에 필요한 권한 부족
- Repository 이름이 잘못됨

**해결 방법:**
1. Token이 유효한지 확인: `curl -H "Authorization: token YOUR_TOKEN" https://api.github.com/user`
2. Token에 `repo`, `admin:repo_hook` 권한이 있는지 확인
3. GITHUB_REPO 형식 확인: `owner/repository` (예: myusername/my-repo)
4. Repository가 실제로 존재하는지 확인

### Connection Refused 오류

**원인:**
- Runner 서비스가 정상 설치되지 않음

**해결 방법:**
```bash
# 서비스 상태 확인
sudo systemctl status github-actions-runner

# 로그 확인
sudo journalctl -u github-actions-runner -f

# 기존 설치 제거 후 다시 시도
cd ~/actions-runner
sudo ./svc.sh uninstall
cd ~
rm -rf ~/actions-runner
./setup-self-hosted-runner.sh
```

## ✅ 설치 확인

1. GitHub 저장소의 Settings → Actions → Runners 확인
2. "dev-server-runner" 가 온라인 상태인지 확인

또는 서버에서:

```bash
sudo systemctl status github-actions-runner
sudo journalctl -u github-actions-runner -f
```

## 🔐 보안 주의사항

- `.env` 파일에 GitHub Token이 포함되므로 **절대 git에 커밋하지 마세요**
- `.gitignore`에 `.env` 파일이 포함되어 있는지 확인하세요
- Token은 주기적으로 rotate 하세요
- 서버는 trusted network에서만 실행하세요
