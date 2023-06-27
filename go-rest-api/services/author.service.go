package services

import (
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/repository"
)

type AuthorServiceInterface interface {
	FindAll() (author []*models.Author, err error)
}

type authorService struct {
	authorRepository repository.AuthorRepositoryInterface
}

func NewAuthorService(authorRepository repository.AuthorRepositoryInterface) AuthorServiceInterface {
	return &authorService{authorRepository}
}

func (s *authorService) FindAll() (author []*models.Author, err error) {
	return s.authorRepository.FindAll()
}
