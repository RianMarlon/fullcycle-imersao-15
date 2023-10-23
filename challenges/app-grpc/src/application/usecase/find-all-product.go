package usecase

import (
	"app-grpc/src/domain/entities"
	"app-grpc/src/domain/repositories"
)

type FindAllProductsUseCase struct {
	ProductsRepository repositories.IProductRepository
}

func (findAllProductsUseCase *FindAllProductsUseCase) Execute() ([]*entities.Product, error) {
	products, err := findAllProductsUseCase.ProductsRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
