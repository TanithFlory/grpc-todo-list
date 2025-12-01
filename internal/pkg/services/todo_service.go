package services

import (
	"context"

	"github.com/TanithFlory/go-todolist-api/internal/pkg/repositories"
	pb "github.com/TanithFlory/go-todolist-api/internal/proto/todo"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.TodoId) (*pb.BoolResponse, error) {
	return &pb.BoolResponse{}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.TodoId) (*pb.BoolResponse, error) {
	return &pb.BoolResponse{}, nil
}

func (s *TodoService) ListTodos(ctx context.Context, _ *emptypb.Empty) (*pb.ListTodoResponse, error) {
	return &pb.ListTodoResponse{}, nil
}
