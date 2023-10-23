package serverGrpc

import (
	"app-grpc/src/application/usecase"
	repositoriesGorm "app-grpc/src/infra/gorm/repositories"
	"app-grpc/src/infra/grpc/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	productsRepository := repositoriesGorm.NewProductRepository(database)
	findAllProductUseCase := &usecase.FindAllProductsUseCase{
		ProductsRepository: productsRepository,
	}
	createProductUseCase := &usecase.CreateProductUseCase{
		ProductsRepository: productsRepository,
	}
	productsGrpcService := NewProductGrpcService(createProductUseCase, findAllProductUseCase)
	pb.RegisterProductServiceServer(grpcServer, productsGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}
}
