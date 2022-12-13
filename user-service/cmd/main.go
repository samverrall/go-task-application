package main

import (
	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/app"
)

func main() {
	log := logger.New("debug")

	a := app.New(log)
	if err := a.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
