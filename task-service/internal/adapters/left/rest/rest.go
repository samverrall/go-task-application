package rest

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samverrall/task-service/internal/service/task"
)

type Rest struct {
	echo    *echo.Echo
	address string
}

func New(address string) *Rest {
	e := echo.New()
	return &Rest{
		echo:    e,
		address: address,
	}
}

func (r *Rest) InitMiddleware() {
	r.echo.Use(middleware.Logger())
	r.echo.Use(middleware.Recover())
}

func (r *Rest) InitHandlers(ctx context.Context, tasksService task.API) {
	newTaskHandler(ctx, r.echo, tasksService)
}

func (r *Rest) Start(ctx context.Context) error {
	if err := r.echo.Start(r.address); err != nil {
		return fmt.Errorf("%w: failed to start server", err)
	}
	return nil
}
