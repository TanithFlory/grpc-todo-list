package jobs

import (
	"log"
	"net"

	"github.com/TanithFlory/go-todolist-api/internal/pkg/services"
	"github.com/TanithFlory/go-todolist-api/internal/proto/todo"
	"google.golang.org/grpc"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")

	grpcServer := grpc.NewServer()
	todoService := services.NewTodoService()
	todo.RegisterTodoServiceServer(grpcServer, todoService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
