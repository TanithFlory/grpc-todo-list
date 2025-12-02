package services

import (
	"context"

	mapper "github.com/TanithFlory/go-todolist-api/internal/pkg/mappers"
	"github.com/TanithFlory/go-todolist-api/internal/pkg/repositories"
	pb "github.com/TanithFlory/go-todolist-api/internal/proto/todo"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.TodoResponse, error) {
	todo, err := s.repo.CreateTodo(ctx, req.Title, req.Description)

	if err != nil {
		return &pb.TodoResponse{}, err
	}

	return &pb.TodoResponse{
		Id:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description.String,
	}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.BoolResponse, error) {
	updated, err := s.repo.UpdateTodo(ctx, req.Id, req.Data.Title, req.Data.Description)

	if err != nil {
		return &pb.BoolResponse{}, err
	}

	return &pb.BoolResponse{
		Success: updated,
	}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.TodoId) (*pb.BoolResponse, error) {
	success, err := s.repo.DeleteTodo(ctx, req.Id)

	if err != nil {
		return &pb.BoolResponse{}, err
	}

	return &pb.BoolResponse{
		Success: success,
	}, nil
}

func (s *TodoService) ListTodos(ctx context.Context, req *pb.Pagination) (*pb.ListTodoResponse, error) {
	offset := req.Offset
	if offset < 0 {
		offset = 0
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	println(req.Limit, req.Offset)

	todos, err := s.repo.ListTodos(ctx, &repositories.Pagination{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	pbTodos := mapper.ToPbTodoList(todos)

	return &pb.ListTodoResponse{
		Todos: pbTodos,
	}, nil
}
