package handler

import "example/gosterx/todo/repository"

type Handler struct {
	TodoRepository repository.TodoRepository
}

func NewHandler(tr repository.TodoRepository) *Handler {
	return &Handler{
		TodoRepository: tr,
	}
}