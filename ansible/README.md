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

`dev` 브랜치에 코드가 push되면 자동으로 배포됩니다.

워크플로우 파일: `.github/workflows/deploy-dev.yml`

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
- SSH 연결 확인
- 서버 디스크 공간 확인
- 권한 확인

### 서비스 시작 실패
- 로그 확인: `journalctl -u tiny-breakers-game`
- 환경 변수 확인: `.env` 파일
- 데이터베이스 연결 확인
