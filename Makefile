lint-task: 
	cd ./task-service && golangci-lint --enable gosec,misspell run ./... &&	go test --race -v ./...

build-task-api:
	go build -o ./task-rest-service github.com/samverrall/task-service/cmd/rest-api

lint-user: 
	cd ./user-service && golangci-lint --enable gosec,misspell run ./... &&	go test --race -v ./...

user-gen-proto:
	cd ./user-service/internal/adapters/left/grpc && protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ./proto/*.proto

build-user:
	go build -o ./user-service-grpc github.com/samverrall/go-task-application/user-service/cmd/grpc
