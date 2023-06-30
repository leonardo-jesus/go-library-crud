package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/repository"
)

const NO_AUTHORS_FOUND_ERROR_MESSAGE = "no authors found"

type AuthorServiceInterface interface {
	FindAll(page int) (author []*models.Author, err error)
	FindByName(name string, page int) (author []*models.Author, err error)
	Create() (err error)
}

type authorService struct {
	authorRepository repository.AuthorRepositoryInterface
}

func NewAuthorService(authorRepository repository.AuthorRepositoryInterface) AuthorServiceInterface {
	return &authorService{authorRepository}
}

func GetCsvReader() (csvReader *csv.Reader, err error) {
	csvFile, err := os.Open("./internal/author/assets/authors.csv")
	if err != nil {
		return nil, fmt.Errorf("os.Open %w", err)
	}
	defer csvFile.Close()

	return csv.NewReader(csvFile), nil
}

func (s *authorService) FindAll(page int) (author []*models.Author, err error) {
	foundAuthors, _ := s.authorRepository.FindAll(page)
	if foundAuthors == nil {
		return nil, errors.New(NO_AUTHORS_FOUND_ERROR_MESSAGE)
	}
	return foundAuthors, nil
}

func (s *authorService) FindByName(name string, page int) (author []*models.Author, err error) {
	foundAuthors, _ := s.authorRepository.FindByName(name, page)
	if foundAuthors == nil {
		return nil, errors.New(NO_AUTHORS_FOUND_ERROR_MESSAGE)
	}
	return foundAuthors, nil
}

func (s *authorService) Create() (err error) {
	csvReader, _ := GetCsvReader()
	s.authorRepository.Create(csvReader)

	return nil
}
