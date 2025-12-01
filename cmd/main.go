package main

import "github.com/TanithFlory/go-todolist-api/internal/app/jobs"

func main() {
	jobs.StartGrpcServer()
}
