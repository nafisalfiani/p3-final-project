.PHONY: swag-stable-install
swag-stable-install:
	@go install github.com/swaggo/swag/cmd/swag@v1.6.7

.PHONY: swag-latest-install
swag-latest-install:
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: swag
swag:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./main.go -o ./docs/swagger --parseInternal

.PHONY: build
build:
	@go build -o ./build/app ./main.go

.PHONY: run
run: swag build
	@./build/app

.PHONY: mock-install
mock-install:
	@go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: mock
mock:
	@`go env GOPATH`/bin/mockgen -source domain/$(domain).go -destination domain/mock/$(domain).go
