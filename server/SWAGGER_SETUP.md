# Swagger 설정 가이드

## Swagger CLI 설치

### 방법 1: go install 사용 (권장)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

설치 후 PATH 확인:
```bash
# Go bin 경로 확인
echo $GOPATH/bin
# 또는
echo $HOME/go/bin
```

PATH에 추가 (필요한 경우):
```bash
# ~/.zshrc 또는 ~/.bashrc에 추가
export PATH=$PATH:$(go env GOPATH)/bin
# 또는
export PATH=$PATH:$HOME/go/bin

# 적용
source ~/.zshrc  # 또는 source ~/.bashrc
```

### 방법 2: Makefile 사용

```bash
make install-swag
```

## Swagger 문서 생성

### 방법 1: 직접 명령어 실행

```bash
cd server
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

### 방법 2: Makefile 사용

```bash
make swagger
```

이 명령어는 자동으로 swag CLI를 설치한 후 문서를 생성합니다.

## 문제 해결

### "command not found: swag" 에러

1. **Go가 설치되어 있는지 확인**:
```bash
go version
```

2. **swag CLI 설치**:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

3. **PATH 확인**:
```bash
# Go bin 경로 확인
go env GOPATH
# 출력 예: /Users/username/go

# 해당 경로의 bin 폴더가 PATH에 있는지 확인
echo $PATH | grep $(go env GOPATH)/bin
```

4. **PATH에 추가** (없는 경우):
```bash
# ~/.zshrc 파일 편집
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

### swag 명령어를 찾을 수 없는 경우

전체 경로로 실행:
```bash
$(go env GOPATH)/bin/swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

또는:
```bash
$HOME/go/bin/swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

## Swagger 문서 확인

문서 생성 후 서버를 실행하고 다음 URL로 접속:
```
http://localhost:8080/swagger/index.html
```
