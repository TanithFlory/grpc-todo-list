package jobs

import (
	"log"
	"net"

	"github.com/TanithFlory/go-todolist-api/internal/config"
	db "github.com/TanithFlory/go-todolist-api/internal/pkg/db"
	"github.com/TanithFlory/go-todolist-api/internal/pkg/repositories"
	"github.com/TanithFlory/go-todolist-api/internal/pkg/services"
	"github.com/TanithFlory/go-todolist-api/internal/proto/todo"
	"google.golang.org/grpc"
)

func StartGrpcServer() {

	grpcServer := grpc.NewServer()
	appConfig, err := config.Load()

	if err != nil {
		log.Printf("Error loading config, %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")

	db, err := db.Connect(appConfig)

	if err != nil {
		log.Fatalf("Unable to connect to db %v", err)
	}

	repository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(repository)
	todo.RegisterTodoServiceServer(grpcServer, todoService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
