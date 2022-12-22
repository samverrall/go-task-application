run: 
	./gateway-service-bin -port 8001 -log-level "debug" -config-path "/Users/samverrall/projects/go-task-application/gateway-service/config/local.config.yaml" && ./user-service-grpc -port 8002


lint-task: 
	cd ./task-service && golangci-lint --enable gosec,misspell run ./... &&	go test --race -v ./...

build-task-api:
	go build -o ./task-rest-service github.com/samverrall/task-service/cmd/rest-api

lint-user: 
	cd ./user-service && golangci-lint --enable gosec,misspell run ./... &&	go test --race -v ./...

lint-gateway: 
	cd ./gateway-service && golangci-lint --enable gosec,misspell run ./... &&	go test --race -v ./...

run-gateway: 
	./gateway-service-bin -port 8001 -log-level "debug" -config-path "/Users/samverrall/projects/go-task-application/gateway-service/config/local.config.yaml"

build-gateway:
	go build -o ./gateway-service-bin github.com/samverrall/go-task-application/gateway-service/cmd

user-gen-proto:
	cd ./user-service/internal/adapters/left/grpc && protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ./proto/*.proto

build-user:
	go build -o ./user-service-grpc github.com/samverrall/go-task-application/user-service/cmd/grpc
