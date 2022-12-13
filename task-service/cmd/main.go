package main

import (
	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/app"
	"github.com/samverrall/task-service/pkg/config"
)

func main() {
	c := config.New()
	log := logger.New("debug")

	a := app.New(log, c)
	if err := a.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
