.PHONY: proto
proto:
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/ticket/ticket.proto
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/category/category.proto
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/region/region.proto
	@protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. handler/grpc/wishlist/wishlist.proto

.PHONY: build
build:
	@go build -o ./build/app ./main.go

.PHONY: run
run: build
	@./build/app
