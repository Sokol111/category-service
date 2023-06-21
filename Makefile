compile-proto:
	protoc -Iproto --go_out=. --go_opt=module=github.com/Sokol111/category-service --go-grpc_out=. --go-grpc_opt=module=github.com/Sokol111/category-service proto/*.proto

generate-wire: compile-proto
	wire ./internal/cmd/server

generate-mocks: generate-wire
	mockery

test: generate-mocks
	go test ./... -v -cover

build-server: test
	go build -o category-service ./internal/cmd

run-server: build-server
	./category-service