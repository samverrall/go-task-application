package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samverrall/task-service/internal/service"
)

func createTask(ctx context.Context, taskService service.TaskServicer) echo.HandlerFunc {
	return func(c echo.Context) error {
		// input its a DTO (Data Transfer Object) that defines a
		// specific REST adapter model to be parsed to a domain.Task.
		// DTOs should use primative types, that can map to Object Value types
		// in the domain.
		var input struct {
			Name string `json:"name"`
		}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		data := service.CreateTaskDTO(input)
		task, err := taskService.CreateTask(ctx, &data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{
				Message: err.Error(),
			})
		}

		output := struct {
			UUID      string `json:"uuid"`
			Name      string `json:"name"`
			CreatedAt string `json:"createdAt"`
		}{
			Name:      task.Name.String(),
			UUID:      task.UUID.String(),
			CreatedAt: task.CreatedAt.Format(time.RFC3339),
		}
		return c.JSON(http.StatusCreated, output)
	}
}

func newTaskHandler(ctx context.Context, e *echo.Echo, taskService service.TaskServicer) {
	e.POST("/api/tasks", createTask(ctx, taskService))
}
