package repository

import (
	"database/sql"
	"errors"
	"example/gosterx/todo/model"
	"log"

	"github.com/google/uuid"
)

// TODO: investigate statement, investigate context - precompile queries
type DatabaseRepository struct {
	dbConn *sql.DB
}

func NewDatabaseRepository(dbConn *sql.DB) *DatabaseRepository {
	return &DatabaseRepository{
		dbConn: dbConn,
	}
}

func (repo *DatabaseRepository) Create(todo *model.Todo) *model.Todo {
	query := "INSERT INTO todos (id, description, status, issued_at) VALUES (?, ?, ?, ?)"
	
	repo.dbConn.Exec(query, todo.Id, todo.Description, todo.Status, todo.IssuedAt)
	
	return todo
}

func (repo *DatabaseRepository) Get(id uuid.UUID) (*model.Todo, error) {
	var todo model.Todo
	
	query := "SELECT id, description, status, issued_at FROM todos WHERE id = ?"

	row := repo.dbConn.QueryRow(query, id)
	if err := row.Scan(&todo.Id, &todo.Description, &todo.Status, &todo.IssuedAt); err != nil {
		if err == sql.ErrNoRows {
			return &todo, errors.New("not found")
		}
		return &todo, errors.New("error")
	}

	return &todo, nil
}

func (repo *DatabaseRepository) GetAll() *[]model.Todo {
	var todos []model.Todo

	query := "SELECT id, description, status, issued_at FROM todos"

	rows, queryErr := repo.dbConn.Query(query)

	if queryErr != nil {
		log.Printf("Error was raised: %s", queryErr.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Status, &todo.IssuedAt); err != nil {
			return nil
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
        return nil
    }

	return &todos
}

func (repo *DatabaseRepository) Update(id uuid.UUID, desctiption string) (*model.Todo, error) {
	query := "UPDATE todos SET description = ? WHERE id = ?"

	repo.dbConn.Exec(query, desctiption, id)

	todo, _ := repo.Get(id)

	return todo, nil
}

func (repo *DatabaseRepository) Delete(id uuid.UUID) error { 
	query := "DELETE FROM todos WHERE id = ?"

	repo.dbConn.Exec(query, id)

	return nil
}