package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	todosGroups := v1.Group("/todos")

	todosGroups.POST("", h.CreateTodo)
	todosGroups.GET("", h.GetAllTodos)
	todosGroups.GET("/:id", h.Get)
	todosGroups.PUT("/:id", h.Update)
	todosGroups.DELETE("/:id", h.Delete)
}