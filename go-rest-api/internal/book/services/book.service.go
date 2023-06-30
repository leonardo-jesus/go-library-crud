package services

import (
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/repository"
)

type BookServiceInterface interface {
	FindAll(page int) (book []*models.BookDB, err error)
	FindByName(name string, page int) (book []*models.BookDB, err error)
	Create(book *models.BookAPI) (err error)
}

type bookService struct {
	bookRepository repository.BookRepositoryInterface
}

func NewBookService(bookRepository repository.BookRepositoryInterface) BookServiceInterface {
	return &bookService{bookRepository}
}

func (s *bookService) FindAll(page int) (book []*models.BookDB, err error) {
	return s.bookRepository.FindAll(page)
}

func (s *bookService) FindByName(name string, page int) (book []*models.BookDB, err error) {
	return s.bookRepository.FindByName(name, page)
}

func (s *bookService) Create(book *models.BookAPI) (err error) {
	s.bookRepository.Create(book)

	return nil
}
