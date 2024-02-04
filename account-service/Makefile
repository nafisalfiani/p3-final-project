.PHONY: proto
proto:
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/user/user.proto
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/auth/auth.proto

.PHONY: build
build:
	@go build -o ./build/app ./main.go

.PHONY: run
run: build
	@./build/app
