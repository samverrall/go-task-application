package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samverrall/task-service/internal/app/api"
	"github.com/samverrall/task-service/internal/domain"
)

func createTask(ctx context.Context, taskService api.TaskServicer) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task domain.Task
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		createdTask, err := taskService.CreateTask(ctx, &task)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, createdTask)
	}
}

func newTaskHandler(ctx context.Context, e *echo.Echo, taskService api.TaskServicer) {
	e.POST("/api/tasks", createTask(ctx, taskService))
}
