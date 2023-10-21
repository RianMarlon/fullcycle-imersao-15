package serverGrpc

import (
	"app-grpc/src/application/usecase"
	"app-grpc/src/infra/grpc/pb"
	"context"
)

type ProductGrpcService struct {
	FindAllProductsUseCase *usecase.FindAllProductsUseCase
	CreateProductUseCase   *usecase.CreateProductUseCase
	pb.UnimplementedProductServiceServer
}

func (productGrpcService *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := productGrpcService.FindAllProductsUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var allProducts []*pb.Product
	for _, product := range products {
		allProducts = append(allProducts, &pb.Product{
			Id:          int32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}
	return &pb.FindProductsResponse{
		Products: allProducts,
	}, nil
}

func (productGrpcService *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productCreated, err := productGrpcService.CreateProductUseCase.Execute(in.Name, in.Description, in.Price)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          int32(productCreated.ID),
			Name:        productCreated.Name,
			Description: productCreated.Description,
			Price:       productCreated.Price,
		},
	}, nil
}

func NewProductGrpcService(createProductUseCase *usecase.CreateProductUseCase, findAllProductsUseCase *usecase.FindAllProductsUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		CreateProductUseCase:   createProductUseCase,
		FindAllProductsUseCase: findAllProductsUseCase,
	}
}
