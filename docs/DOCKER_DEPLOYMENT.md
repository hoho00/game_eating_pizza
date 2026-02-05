# Docker 배포 가이드

## 📋 개요

이 프로젝트는 이제 Docker와 docker-compose를 사용하여 배포됩니다.

## 🏗️ 아키텍처

```
┌─────────────────────────────────────────┐
│      GitHub Actions (Self-hosted)       │
├─────────────────────────────────────────┤
│ 1. Checkout code                        │
│ 2. Copy .env file                       │
│ 3. Run Ansible playbook                 │
└─────────────┬───────────────────────────┘
              │
              ▼
┌─────────────────────────────────────────┐
│      Ansible (deploy.yml)               │
├─────────────────────────────────────────┤
│ 1. Check docker-compose.yml             │
│ 2. Stop existing containers             │
│ 3. Build Docker images                  │
│ 4. Start containers                     │
│ 5. Health check                         │
│ 6. Display logs                         │
└─────────────┬───────────────────────────┘
              │
              ▼
┌─────────────────────────────────────────┐
│      Docker Containers                  │
├─────────────────────────────────────────┤
│ Game Server        │  Batch Server       │
│ Port: 8080         │  (No exposed port)  │
│ Image: ...game     │  Image: ...batch    │
└─────────────────────────────────────────┘
```

## 🐳 Docker 이미지 구조

### Game Server (Dockerfile)
- Multi-stage build
- Go 1.23 기반
- Alpine Linux (최소화)
- Non-root 사용자 실행
- Health check 포함

### Batch Server (Dockerfile.batch)
- Multi-stage build
- Go 1.23 기반
- Alpine Linux (최소화)
- Non-root 사용자 실행

## 📦 Docker Compose 설정

```yaml
services:
  game-server:
    image: tiny-breakers-game:latest
    ports:
      - "8080:8080"
    
  batch-server:
    image: tiny-breakers-batch:latest
    depends_on:
      - game-server
```

## 🚀 배포 프로세스

### 1. Self-hosted Runner에서 (자동)

```bash
# GitHub Actions이 자동 실행:
1. 코드 체크아웃
2. .env 파일 복사 (존재하면)
3. Go 의존성 검증
4. Ansible 플레이북 실행
```

### 2. Ansible에서 (자동)

```bash
cd /home/ki/src/game/server
docker-compose down        # 기존 컨테이너 종료
docker-compose build       # 이미지 빌드
docker-compose up -d       # 컨테이너 시작
```

### 3. 수동 배포

```bash
# 로컬 서버에서 직접 실행
cd /home/ki/src/game/server

# 컨테이너 빌드 및 시작
docker-compose up -d

# 로그 확인
docker-compose logs -f game-server

# 컨테이너 중지
docker-compose down
```

## 🔍 모니터링 및 관리

### 컨테이너 상태 확인
```bash
docker ps --filter "name=tiny-breakers"
```

### 로그 확인
```bash
# Game Server 로그
docker logs -f tiny-breakers-game-server

# Batch Server 로그
docker logs -f tiny-breakers-batch-server

# 최근 20줄만 보기
docker logs --tail 20 tiny-breakers-game-server
```

### 컨테이너 접근
```bash
docker exec -it tiny-breakers-game-server /bin/sh
```

### 컨테이너 재시작
```bash
docker-compose -f /home/ki/src/game/server/docker-compose.yml restart
```

### 완전 재배포
```bash
docker-compose -f /home/ki/src/game/server/docker-compose.yml down -v
docker-compose -f /home/ki/src/game/server/docker-compose.yml up -d --build
```

## 📝 환경 변수 설정

`.env` 파일이 server 디렉토리에 있어야 합니다:

```bash
# /home/ki/src/game/.env (자동으로 server 디렉토리로 복사됨)
DATABASE_URL=your_database_url
REDIS_URL=your_redis_url
JWT_SECRET=your_secret_key
LOG_LEVEL=info
```

## 🐛 문제 해결

### 이미지 빌드 실패
```bash
# go.mod 파일 확인
ls -la /home/ki/src/game/server/go.mod

# 수동으로 빌드 재시도
cd /home/ki/src/game/server
docker-compose build --no-cache
```

### 컨테이너 실행 실패
```bash
# 컨테이너 로그 확인
docker logs tiny-breakers-game-server

# 존재하는 이미지 확인
docker images | grep tiny-breakers

# 존재하는 컨테이너 확인
docker ps -a | grep tiny-breakers
```

### 포트 충돌
```bash
# 포트 확인
netstat -tlnp | grep 8080

# 기존 프로세스 종료 후 재배포
docker-compose down
docker-compose up -d
```

## 📚 관련 파일

- [Dockerfile](../../server/Dockerfile) - Game Server 이미지
- [Dockerfile.batch](../../server/Dockerfile.batch) - Batch Server 이미지
- [docker-compose.yml](../../server/docker-compose.yml) - 컨테이너 정의
- [deploy.yml](../playbooks/deploy.yml) - Ansible 배포 플레이북
- [.github/workflows/deploy-dev.yml](../../.github/workflows/deploy-dev.yml) - CI/CD 설정

## ✅ 배포 완료 확인

```bash
# 1. 컨테이너 실행 확인
docker ps --filter "name=tiny-breakers"

# 2. 로그에서 정상 시작 확인
docker logs tiny-breakers-game-server

# 3. Health check 확인
curl http://localhost:8080/health

# 4. Ansible 플레이북 출력에서 ✅ Deployment completed! 확인
```
