package entities

import (
	"github.com/asaskevich/govalidator"
)

type Product struct {
	ID          int     `valid:"-" gorm:"column:id;int;primaryKey;autoIncrement"`
	Name        string  `valid:"required" gorm:"column:name;type:varchar(255);not null"`
	Description string  `valid:"required" gorm:"column:description;type:text;not null"`
	Price       float32 `valid:"required" gorm:"column:price;type:float;not null"`
}

func (product *Product) IsValid() error {
	_, err := govalidator.ValidateStruct(product)

	if err != nil {
		return err
	}

	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	err := product.IsValid()
	if err != nil {
		return nil, err
	}
	return &product, nil
}
