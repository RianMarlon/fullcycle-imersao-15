package usecase

import (
	"app-grpc/src/domain/entities"
	"app-grpc/src/domain/repositories"
)

type CreateProductUseCase struct {
	ProductsRepository repositories.IProductRepository
}

func (createProductUseCase *CreateProductUseCase) Execute(name string, description string, price float32) (*entities.Product, error) {
	product, err := entities.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}
	productCreated, err := createProductUseCase.ProductsRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return productCreated, nil
}
