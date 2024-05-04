package conn

import (
	"github.com/Asad2730/BookApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnRepository struct {
	db *gorm.DB
}

func NewConnection() (*PostgresConnRepository, error) {
	dsn := "host=localhost user=postgres password=123 dbname=book port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.Book{})
	return &PostgresConnRepository{db: db}, nil
}
