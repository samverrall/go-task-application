package rest

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samverrall/task-service/internal/port/service/task"
)

type Rest struct {
	echo         *echo.Echo
	address      string
	tasksService task.API
}

func New(address string, taskService task.API) *Rest {
	e := echo.New()
	return &Rest{
		echo:         e,
		address:      address,
		tasksService: taskService,
	}
}

func (r *Rest) Start(ctx context.Context) error {
	r.echo.Use(middleware.Logger())
	r.echo.Use(middleware.Recover())

	// Register API Handlers
	r.newTaskHandler(ctx)

	if err := r.echo.Start(r.address); err != nil {
		return fmt.Errorf("%w: failed to start server", err)
	}
	return nil
}
