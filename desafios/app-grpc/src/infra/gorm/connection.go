package database

import (
	"app-grpc/src/domain/entities"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectionDB() *gorm.DB {
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		err = db.AutoMigrate(&entities.Product{})

		if err != nil {
			log.Fatalf("Error running migrations: %v", err)
			panic(err)
		}
	}

	return db
}
