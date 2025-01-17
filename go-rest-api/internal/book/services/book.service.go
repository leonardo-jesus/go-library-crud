package services

import (
	"errors"
	"fmt"

	"dario.cat/mergo"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/repository"
)

const NO_BOOKS_FOUND_ERROR_MESSAGE = "no books found"

type BookServiceInterface interface {
	FindAll(page int) (book []*models.Book, err error)
	FindByFilteredBooks(filters models.FilteredBookSchema, page int) (book []*models.Book, err error)
	FindById(id int) (book *models.UpdateBookSchema, err error)
	Create(book *models.CreateBookSchema) (createdBook *models.CreateBookSchema, err error)
	Update(book *models.UpdateBookSchema) (updatedBook *models.UpdateBookSchema, err error)
	Delete(id int) (err error)
}

type bookService struct {
	bookRepository repository.BookRepositoryInterface
}

func NewBookService(bookRepository repository.BookRepositoryInterface) BookServiceInterface {
	return &bookService{bookRepository}
}

func (s *bookService) FindAll(page int) (book []*models.Book, err error) {
	foundBooks, _ := s.bookRepository.FindAll(page)
	if foundBooks == nil {
		return nil, errors.New(NO_BOOKS_FOUND_ERROR_MESSAGE)
	}
	return foundBooks, nil
}

func (s *bookService) FindByFilteredBooks(filters models.FilteredBookSchema, page int) (book []*models.Book, err error) {
	foundBooks, _ := s.bookRepository.FindByFilteredBooks(filters, page)
	if foundBooks == nil {
		return nil, errors.New(NO_BOOKS_FOUND_ERROR_MESSAGE)
	}
	return foundBooks, nil
}

func (s *bookService) FindById(id int) (book *models.UpdateBookSchema, err error) {
	foundBook, _ := s.bookRepository.FindById(id)
	if foundBook == nil {
		return nil, errors.New(NO_BOOKS_FOUND_ERROR_MESSAGE)
	}
	return foundBook, nil
}

func (s *bookService) Create(book *models.CreateBookSchema) (createdBook *models.CreateBookSchema, err error) {
	err = s.bookRepository.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) Update(book *models.UpdateBookSchema) (updatedBook *models.UpdateBookSchema, err error) {
	bookFromDb, err := s.FindById(book.Id)
	if err != nil {
		return nil, err
	}

	err = mergo.Merge(bookFromDb, book, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("error merging structs at update service: %w", err)
	}

	err = s.bookRepository.Update(bookFromDb)
	if err != nil {
		return nil, err
	}

	return bookFromDb, nil
}

func (s *bookService) Delete(id int) (err error) {
	return s.bookRepository.Delete(id)
}
