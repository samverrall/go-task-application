build-task-api:
	go build  -o ./task-rest-service github.com/samverrall/task-service/cmd/rest-api

test-task:
	cd ./task-service && go test --race -v ./...

build-user:
	go build  -o ./user-service-bin github.com/samverrall/go-task-application/user-service/cmd
