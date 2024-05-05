package interfaces

import (
	"github.com/Asad2730/BookApi/models"
)

type BookRepo interface {
	Create(book *models.Book) error
	ReadAll() (*[]models.Book, error)
	Read(id int) (*models.Book, error)
	Update(id int, book *models.Book) error
	Delete(id int) error
}
