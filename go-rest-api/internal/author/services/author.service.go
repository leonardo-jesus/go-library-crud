package services

import (
	"encoding/csv"
	"os"

	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/models"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/repository"
)

type AuthorServiceInterface interface {
	FindAll() (author []*models.Author, err error)
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

func (s *authorService) FindAll() (author []*models.Author, err error) {
	return s.authorRepository.FindAll()
}

func (s *authorService) Create() (err error) {
	// records := ReadCsv("./internal/assets/authors.csv")

	// var authorRecords []models.Author
	// for _, record := range records {
	// 	data := models.Author{
	// 		Name: record,
	// 	}

	// 	// TODO: For each record, save in the DB (should make a chunk and save all together on DB)
	// 	authorRecords = append(authorRecords, data)
	// }

	// fmt.Println(authorRecords)

	s.authorRepository.Create()

	return nil
}
