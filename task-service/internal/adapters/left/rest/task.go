package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samverrall/task-service/internal/port/service/task"
)

func (r *Rest) createTask(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		// input its a DTO (Data Transfer Object) that defines a
		// specific REST adapter model to be parsed to a domain.Task.
		// DTOs should use primitive types, that can map to Object Value types
		// in the domain.
		var input struct {
			Name       string    `json:"name"`
			CompleteBy time.Time `json:"completeBy"`
		}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		data := task.CreateTaskDTO(input)
		task, err := r.tasksService.CreateTask(ctx, data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{
				Message: err.Error(),
			})
		}

		output := struct {
			UUID       string `json:"uuid"`
			Name       string `json:"name"`
			CompleteBy string `json:"completeBy"`
			CreatedAt  string `json:"createdAt"`
		}{
			Name:       task.Name.String(),
			UUID:       task.UUID.String(),
			CompleteBy: task.CompleteBy.Time().Format(time.RFC3339),
			CreatedAt:  task.CreatedAt.Format(time.RFC3339),
		}
		return c.JSON(http.StatusCreated, output)
	}
}

func (r *Rest) newTaskHandler(ctx context.Context) {
	r.echo.POST("/api/tasks", r.createTask(ctx))
}
