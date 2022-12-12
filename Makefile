build-task:
	go build --race -o ./task-service-bin github.com/samverrall/task-service/cmd

build-user:
	go build --race -o ./user-service-bin github.com/samverrall/go-task-application/user-service/cmd
