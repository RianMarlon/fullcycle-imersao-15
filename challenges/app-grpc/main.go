package main

import (
	database "app-grpc/src/infra/gorm"
	serverGrpc "app-grpc/src/infra/grpc"
)

func main() {
	database := database.ConnectionDB()
	serverGrpc.StartGrpcServer(database, 50051)
}
