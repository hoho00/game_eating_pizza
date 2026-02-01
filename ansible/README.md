# Ansible 배포 설정

Tiny Breakers 프로젝트의 Ansible 배포 설정입니다.

## 구조

```
ansible/
├── inventory/
│   └── dev.yml          # 개발 서버 인벤토리
├── playbooks/
│   └── deploy.yml       # 배포 Playbook
├── templates/
│   ├── game-server.service.j2    # 게임 서버 systemd 서비스 템플릿
│   └── batch-server.service.j2   # 배치 서버 systemd 서비스 템플릿
├── ansible.cfg          # Ansible 설정
└── README.md            # 이 파일
```

## 배포 전략

### One Codebase, Multiple Deployments

같은 코드베이스에서 두 개의 서비스를 배포합니다:

1. **게임 서버** (`~/backend/game-server`)
   - HTTP API 서버
   - 진입점: `cmd/server/main.go`

2. **배치 스케줄링 서버** (`~/backend/batch-server`)
   - 배치 작업 실행
   - 진입점: `cmd/batch/main.go` (추후 생성 필요)

## 수동 배포

### 사전 요구사항

- Ansible 2.9 이상
- Python 3.7 이상
- Go 1.23 이상 (로컬 빌드용)
- SSH 접근 권한 (개발 서버)

### 배포 실행

1. **로컬에서 바이너리 빌드**:
```bash
cd server
go mod download
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server ./cmd/server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server-batch ./cmd/batch
```

2. **Ansible로 배포**:
```bash
cd ansible
ansible-playbook playbooks/deploy.yml -i inventory/dev.yml
```

**참고**: 바이너리는 `server/` 디렉토리에 있어야 합니다.

## 자동 배포 (GitHub Actions)

### 방법 1: Self-hosted Runner 사용 (권장) ⭐

개발 서버에 GitHub Actions Self-hosted Runner를 설치하면 SSH 연결 문제를 해결할 수 있습니다.

#### Self-hosted Runner 설치

1. **GitHub Personal Access Token 생성**:
   - GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
   - `repo` 권한 필요

2. **개발 서버에서 Runner 설치**:
```bash
# 스크립트를 개발 서버로 전송
scp ansible/scripts/setup-self-hosted-runner.sh ki@100.69.72.59:~/

# 개발 서버에 SSH 접속
ssh ki@100.69.72.59

# 스크립트 실행
chmod +x setup-self-hosted-runner.sh
./setup-self-hosted-runner.sh
```

3. **워크플로우 파일 사용**:
   - `.github/workflows/deploy-dev-self-hosted.yml` 사용
   - 또는 기존 파일에서 `runs-on: ubuntu-latest` → `runs-on: self-hosted`로 변경

#### 장점
- ✅ SSH 연결 불필요
- ✅ 방화벽 설정 불필요
- ✅ 빠른 배포 (로컬 네트워크)

### 방법 2: GitHub-hosted Runner 사용 (SSH 연결 필요)

**주의**: 개발 서버의 방화벽에서 GitHub Actions IP를 허용해야 합니다.

워크플로우 파일: `.github/workflows/deploy-dev.yml`

**SSH 연결 문제 해결 방법**:
1. 방화벽에서 포트 22(SSH) 허용
2. GitHub Actions IP 범위 허용: https://api.github.com/meta (IP 주소 확인)
3. 또는 SSH 포트 변경 (인벤토리 파일에서 `ansible_port` 설정)

## 서버 정보

- **호스트**: 100.69.72.59
- **사용자**: ki
- **비밀번호**: 4939 (인벤토리 파일에 저장됨)

**보안 주의**: 프로덕션 환경에서는 SSH 키 기반 인증을 사용하고, 비밀번호를 인벤토리 파일에 저장하지 마세요.

## 배포 프로세스

### GitHub Actions 자동 배포
1. `dev` 브랜치에 push
2. GitHub Actions에서 코드 체크아웃
3. **GitHub Actions에서 Go 빌드** (게임 서버 + 배치 서버)
4. Ansible을 통해 빌드된 바이너리를 개발 서버로 전송
5. systemd 서비스 재시작

**중요**: 개발 서버는 GitHub에 접근할 필요가 없습니다. 모든 빌드는 GitHub Actions에서 수행됩니다.

## 서비스 관리

배포 후 서비스는 systemd로 관리됩니다:

```bash
# 게임 서버 상태 확인
sudo systemctl status tiny-breakers-game

# 배치 서버 상태 확인
sudo systemctl status tiny-breakers-batch

# 로그 확인
sudo journalctl -u tiny-breakers-game -f
sudo journalctl -u tiny-breakers-batch -f
```

## 문제 해결

### 빌드 실패
- Go 버전 확인: `go version`
- 의존성 확인: `go mod download`

### 배포 실패

#### SSH 연결 타임아웃
- **원인**: 개발 서버가 외부에서 접근 불가능 (방화벽, 보안 그룹)
- **해결 방법**:
  1. **Self-hosted Runner 사용** (권장): 위의 "방법 1" 참고
  2. 방화벽에서 포트 22 허용
  3. GitHub Actions IP 범위 허용
  4. SSH 포트 확인 및 변경

#### 기타 문제
- SSH 연결 확인: `ssh ki@100.69.72.59`
- 서버 디스크 공간 확인: `df -h`
- 권한 확인: `ls -la ~/backend/`

### 서비스 시작 실패
- 로그 확인: `journalctl -u tiny-breakers-game`
- 환경 변수 확인: `.env` 파일
- 데이터베이스 연결 확인
