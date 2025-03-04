package services

import (
	"github.com/wisaitas/todo-web/internal/repositories"
	"github.com/wisaitas/todo-web/internal/utils"
)

type TodoService interface {
}

type todoService struct {
	todoRepository repositories.TodoRepository
	redisUtil      utils.RedisClient
}

func NewTodoService(
	todoRepository repositories.TodoRepository,
	redisUtil utils.RedisClient,
) TodoService {
	return &todoService{
		todoRepository: todoRepository,
		redisUtil:      redisUtil,
	}
}
