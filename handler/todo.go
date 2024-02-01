package handler

import (
	"example/gosterx/todo/model"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateTodo(c echo.Context) error {
	t := new(createTodo)
	if err := c.Bind(t); err != nil {
		log.Printf("Error was raised during body decoding: %s", err.Error())
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	todo := model.New(t.Description)

	created := h.TodoRepository.Create(todo)

	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) GetAllTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, h.TodoRepository.GetAll())
}

func (h *Handler) Get(c echo.Context) error {
	id := c.Param("id")
	uuid, uuidErr := uuid.Parse(id)
	if uuidErr != nil {
		log.Printf("Error was raised during id parsing: %s", uuidErr.Error())
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	todo, todoErr := h.TodoRepository.Get(uuid)
	if todoErr != nil {
		log.Printf("Error was raised during todo receiving: %s", todoErr.Error())
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *Handler) Update(c echo.Context) error {
	id := c.Param("id")
	uuid, uuidErr := uuid.Parse(id)
	if uuidErr != nil {
		log.Printf("Error was raised during id parsing: %s", uuidErr.Error())
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	t := new(updateTodo)
	if reqErr := c.Bind(t); reqErr != nil {
		log.Printf("Error was raised during body decoding: %s", reqErr.Error())
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	updated, updateErr := h.TodoRepository.Update(uuid, t.Description)
	if updateErr != nil {
		log.Printf("Error was raised during todo updating: %s", updateErr.Error())
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")
	uuid, uuidErr := uuid.Parse(id)
	if uuidErr != nil {
		log.Printf("Error was raised during id parsing: %s", uuidErr.Error())
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	todoErr := h.TodoRepository.Delete(uuid)
	if todoErr != nil {
		log.Printf("Error was raised during todo deleting: %s", todoErr.Error())
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}