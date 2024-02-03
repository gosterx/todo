package repository

import (
	"example/gosterx/todo/model"

	"github.com/google/uuid"
)

// TODO: add error for every (because of the db)
type TodoRepository interface {
	Create(todo *model.Todo) *model.Todo
	Get(id uuid.UUID) (*model.Todo, error)
	GetAll() *[]model.Todo
	Update(id uuid.UUID, desctiption string) (*model.Todo, error)
	Delete(id uuid.UUID) error
}
