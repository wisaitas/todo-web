package handlers

import "github.com/wisaitas/todo-web/internal/services"

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoService(
	todoService services.TodoService,
) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}
