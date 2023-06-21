compile-proto:
	protoc -Iproto --go_out=. --go_opt=module=github.com/Sokol111/category-service --go-grpc_out=. --go-grpc_opt=module=github.com/Sokol111/category-service proto/*.proto

generate-wire: compile-proto
	@echo "Generating wire files"
	wire ./internal/cmd/server

build-server: generate-wire
	@echo "Building the binary..."
	go build -o category-service ./internal/cmd
	@echo "Done!"

run-server: build-server
	@echo "Starting the server"
	#go run ./internal/cmd
	./category-service