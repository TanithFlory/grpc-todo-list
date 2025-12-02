package mapper

import (
	"github.com/TanithFlory/go-todolist-api/internal/pkg/db/sqlc"
	"github.com/TanithFlory/go-todolist-api/internal/proto/todo"
)

func ToPbTodo(t sqlc.Todo) *todo.TodoResponse {
	return &todo.TodoResponse{Id: t.ID, Title: t.Title, Description: t.Description.String}
}

func ToPbTodoList(list []sqlc.Todo) []*todo.TodoResponse {
	res := make([]*todo.TodoResponse, 0, len(list))
	for _, t := range list {
		res = append(res, ToPbTodo(t))
	}
	return res
}
