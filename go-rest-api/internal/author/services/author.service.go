package services

import (
	"encoding/csv"
	"os"

	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/repository"
)

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

func ReadCsv(filename string) (records []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	records, err = csvReader.Read()
	if err != nil {
		panic(err)
	}

	defer file.Close()
	return records
}

func (s *authorService) FindAll(page int) (author []*models.Author, err error) {
	return s.authorRepository.FindAll(page)
}

func (s *authorService) FindByName(name string, page int) (author []*models.Author, err error) {
	return s.authorRepository.FindByName(name, page)
}

func (s *authorService) Create() (err error) {
	s.authorRepository.Create()

	return nil
}
