package main

import (
	"example/gosterx/todo/handler"
	"example/gosterx/todo/repository"
	"example/gosterx/todo/router"
)

func main() {
	r := router.New()

	v1 := r.Group("/api")

	tr := repository.NewInMemoryRepository()
	h := handler.NewHandler(tr)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":1323"))
}
