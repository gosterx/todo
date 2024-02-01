package handler

type createTodo struct {
	Description string `json:"description"`
}

type updateTodo struct {
	Description string `json:"description"`
}