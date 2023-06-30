package services

import (
	"dario.cat/mergo"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/repository"
)

type BookServiceInterface interface {
	FindAll(page int) (book []*models.Book, err error)
	FindByName(name string, page int) (book []*models.Book, err error)
	FindById(id int) (book *models.UpdateBookSchema, err error)
	Create(book *models.CreateBookSchema) (err error)
	Update(book *models.UpdateBookSchema) (err error)
	Delete(id int) (err error)
}

type bookService struct {
	bookRepository repository.BookRepositoryInterface
}

func NewBookService(bookRepository repository.BookRepositoryInterface) BookServiceInterface {
	return &bookService{bookRepository}
}

func (s *bookService) FindAll(page int) (book []*models.Book, err error) {
	return s.bookRepository.FindAll(page)
}

func (s *bookService) FindByName(name string, page int) (book []*models.Book, err error) {
	return s.bookRepository.FindByName(name, page)
}

func (s *bookService) FindById(id int) (book *models.UpdateBookSchema, err error) {
	return s.bookRepository.FindById(id)
}

func (s *bookService) Create(book *models.CreateBookSchema) (err error) {
	return s.bookRepository.Create(book)
}

func (s *bookService) Update(book *models.UpdateBookSchema) (err error) {
	bookFromDb, err := s.FindById(book.Id)
	if err != nil {
		return err
	}

	err = mergo.Merge(bookFromDb, book, mergo.WithOverride)

	if err != nil {
		return err
	}

	s.bookRepository.Update(bookFromDb)

	return nil
}

func (s *bookService) Delete(id int) (err error) {
	s.bookRepository.Delete(id)

	return nil
}
