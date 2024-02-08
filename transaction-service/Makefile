.PHONY: proto
proto:
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/transaction/transaction.proto
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/wallet/wallet.proto

.PHONY: build
build:
	@go build -o ./build/app ./main.go

.PHONY: run
run: build
	@./build/app
