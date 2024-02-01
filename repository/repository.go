package repository

import (
	"errors"
	"example/gosterx/todo/model"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Create(todo *model.Todo) *model.Todo
	Get(id uuid.UUID) (*model.Todo, error)
	GetAll() *[]model.Todo
	Update(id uuid.UUID, desctiption string) (*model.Todo, error)
	Delete(id uuid.UUID) error
}

type InMemoryRepository struct {
	Todos *[]model.Todo
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Todos: &[]model.Todo{},
	}
}

func (repo *InMemoryRepository) Create(todo *model.Todo) *model.Todo {
	*repo.Todos = append(*repo.Todos, *todo)
	return todo
}

func (repo *InMemoryRepository) Get(id uuid.UUID) (*model.Todo, error) {
	for i := 0; i < len(*repo.Todos); i++ {
		if (*repo.Todos)[i].Id == id {
			return &(*repo.Todos)[i], nil
		}
	}

	return nil, errors.New("not found")
}

func (repo *InMemoryRepository) GetAll() *[]model.Todo {
	return repo.Todos
}

func (repo *InMemoryRepository) Update(id uuid.UUID, desctiption string) (*model.Todo, error) {
	index := -1

	for i := 0; i < len(*repo.Todos); i++ {
		if (*repo.Todos)[i].Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, errors.New("not found")
	}

	(*repo.Todos)[index].Description = desctiption
	return &(*repo.Todos)[index], nil
}

func (repo *InMemoryRepository) Delete(id uuid.UUID) error {
	index := -1

	for i := 0; i < len(*repo.Todos); i++ {
		if (*repo.Todos)[i].Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("not found")
	}

	*(repo.Todos) = append((*repo.Todos)[:index], (*repo.Todos)[index + 1:]...)

	return nil
}