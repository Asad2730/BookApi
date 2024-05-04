package repositories

import (
	"github.com/Asad2730/BookApi/models"
)

type BookRepositories interface {
	Create(book *models.Book) error
	ReadAll() (*[]models.Book, error)
	Read(id int) (*models.Book, error)
	Update(id int, book *models.Book) error
	Delete(id int) error
}

type DbRepo struct {
	repo BookRepositories
}

func NewDbRepo(repo BookRepositories) *DbRepo {
	return &DbRepo{repo: repo}
}

func (r *DbRepo) Create(book *models.Book) error {
	return r.repo.Create(book)
}

func (r *DbRepo) Read(id int) (*models.Book, error) {
	return r.repo.Read(id)
}

func (r *DbRepo) ReadAll() (*[]models.Book, error) {
	return r.repo.ReadAll()
}

func (r *DbRepo) Update(id int, book *models.Book) error {
	return r.repo.Update(id, book)
}

func (r *DbRepo) Delete(id int) error {
	return r.repo.Delete(id)
}
