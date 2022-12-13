package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/samverrall/go-task-application/logger"
)

type App struct {
	logger logger.Logger
}

func New(log logger.Logger) *App {
	return &App{
		logger: log,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		a.logger.Error("signal.Notify: %v", v)
	case done := <-ctx.Done():
		a.logger.Error("ctx.Done: %v", done)
	}

	return nil
}
