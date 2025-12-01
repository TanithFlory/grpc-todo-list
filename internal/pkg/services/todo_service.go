package services

import (
	"context"

	pb "github.com/TanithFlory/go-todolist-api/internal/proto/todo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.TodoResponse, error) {

	return &pb.TodoResponse{}, nil
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
