package main

import (
	"github.com/samverrall/task-service/internal/app"
	"github.com/samverrall/task-service/pkg/config"
	"github.com/samverrall/task-service/pkg/logger"
)

func main() {
	c := config.New()
	log := logger.New("debug")

	a := app.New(log, c)
	if err := a.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
