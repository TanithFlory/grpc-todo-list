package repositories

import (
	"context"
	"database/sql"

	"github.com/TanithFlory/go-todolist-api/internal/pkg/db/sqlc"
	"github.com/google/uuid"
)

type TodoRepository struct {
	q *sqlc.Queries
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		q: sqlc.New(db),
	}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, title string, description string) (sqlc.Todo, error) {
	id := uuid.New().String()

	params := sqlc.CreateTodoParams{
		ID:    id,
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	}

	err := r.q.CreateTodo(ctx, params)

	if err != nil {
		return sqlc.Todo{}, err
	}

	if err != nil {
		return sqlc.Todo{}, err
	}

	return r.q.GetTodo(ctx, id)
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, id string, title string, description string) (bool, error) {
	err := r.q.UpdateTodo(ctx, sqlc.UpdateTodoParams{
		ID:    id,
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id string) (bool, error) {
	err := r.q.DeleteTodo(ctx, id)

	if err != nil {
		return false, err
	}

	return true, nil
}

type Pagination struct {
	Limit  int32
	Offset int32
}

func (r *TodoRepository) ListTodos(ctx context.Context, pagination *Pagination) ([]sqlc.Todo, error) {
	todos, err := r.q.ListTodos(ctx, sqlc.ListTodosParams{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}
