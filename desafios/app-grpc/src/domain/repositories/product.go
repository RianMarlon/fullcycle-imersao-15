package repositories

import "app-grpc/src/domain/entities"

type IProductRepository interface {
	FindAll() ([]*entities.Product, error)
	Create(product *entities.Product) (*entities.Product, error)
}
