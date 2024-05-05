package repositories

import (
	"errors"

	"github.com/Asad2730/BookApi/models"
)

type CustomErr struct {
	message string
}

type InMemoryBookRepository struct {
	books map[uint]*models.Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: make(map[uint]*models.Book),
	}
}

func (r *InMemoryBookRepository) Create(book *models.Book) error {
	r.books[book.ID] = book
	return nil
}

func (r *InMemoryBookRepository) Read(id uint) (*models.Book, error) {
	book, err := r.books[id]
	if err {
		return nil, errors.New("Error get book record")
	}
	return book, nil
}

func (r *InMemoryBookRepository) ReadAll() ([]*models.Book, error) {
	books := make([]*models.Book, 0, len(r.books))

	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func (r *InMemoryBookRepository) Update(id uint, book *models.Book) error {
	_, err := r.books[id]
	if err {
		return errors.New("Error updating")
	}

	r.books[id] = book
	return nil
}

func (r *InMemoryBookRepository) Delete(id uint) error {
	_, err := r.books[id]
	if err {
		return errors.New("Error deleting")
	}
	delete(r.books, id)
	return nil
}
