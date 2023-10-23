package factory

import (
	"codepix-go/application/usecase"
	"codepix-go/infrastructure/repository"

	"github.com/jinzhu/gorm"
)

func TransactonUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{
		Db: database,
	}
	transactionRepository := repository.TransactionRepositoryDb{
		Db: database,
	}
	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         &pixRepository,
	}

	return transactionUseCase
}
