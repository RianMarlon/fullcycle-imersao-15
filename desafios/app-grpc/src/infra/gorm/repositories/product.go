package repositoriesGorm

import (
	"app-grpc/src/domain/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func (productRepository *ProductRepository) FindAll() ([]*entities.Product, error) {
	var products []*entities.Product
	result := productRepository.Db.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (productRepository *ProductRepository) Create(product *entities.Product) (*entities.Product, error) {
	result := productRepository.Db.Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}
