.PHONY: swagger
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
	@echo "Swagger documentation generated successfully!"

.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: build
build:
	@go build -o bin/server cmd/server/main.go

.PHONY: tidy
tidy:
	@go mod tidy
