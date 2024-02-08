.PHONY: proto
proto:
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/email/email.proto

.PHONY: build
build:
	@go build -o ./build/app ./main.go

.PHONY: run
run: build
	@./build/app
