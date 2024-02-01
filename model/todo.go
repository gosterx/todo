package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID
	Description string
	Status      string
	IssuedAt    time.Time
}

func New(description string) *Todo {
	uuid := uuid.New()
	issuedAt := time.Now()
	status := "InProgress"

	todo := Todo {
		Id: uuid,
		Description: description,
		Status: status,
		IssuedAt: issuedAt,
	}
	return &todo;
}