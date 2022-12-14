package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samverrall/task-service/internal/service"
)

// taskPayload its a DTO (Data Transfer Object) that defines a
// specific REST adapter model to be parsed to a domain.Task.
// DTOs should use primative types, that can map to Object Value types
// in the domain.
type taskPayload struct {
	Name string `json:"name"`
}

func createTask(ctx context.Context, taskService service.TaskServicer) echo.HandlerFunc {
	return func(c echo.Context) error {
		var taskPaylaod taskPayload
		if err := c.Bind(&taskPaylaod); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		data := service.CreateTaskDTO(taskPaylaod)
		createdTask, err := taskService.CreateTask(ctx, &data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, createdTask)
	}
}

func newTaskHandler(ctx context.Context, e *echo.Echo, taskService service.TaskServicer) {
	e.POST("/api/tasks", createTask(ctx, taskService))
}
