.PHONY: swagger
swagger:
	@echo "Installing swag CLI..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
	@echo "Swagger documentation generated successfully!"

.PHONY: install-swag
install-swag:
	@echo "Installing swag CLI..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Swagger CLI installed successfully!"
	@echo "Make sure $GOPATH/bin or $HOME/go/bin is in your PATH"

.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: build
build:
	@go build -o bin/server cmd/server/main.go

.PHONY: tidy
tidy:
	@go mod tidy
