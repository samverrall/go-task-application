lint-task: 
	cd ./task-service && golangci-lint --enable gosec,misspell run ./... &&	go test --race -v ./...

build-task-api:
	go build  -o ./task-rest-service github.com/samverrall/task-service/cmd/rest-api


build-user:
	go build  -o ./user-service-bin github.com/samverrall/go-task-application/user-service/cmd
