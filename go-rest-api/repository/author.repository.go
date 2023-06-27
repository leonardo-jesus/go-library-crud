package repository

import (
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/models"
)

const AuthorsCollection = "authors"

type AuthorRepositoryInterface interface {
	FindAll() (author []*models.Author, err error)
}

type authorRepository struct {
	authors []*models.Author
}

func NewAuthorRepository(authors []*models.Author) AuthorRepositoryInterface {
	return &authorRepository{authors}
}

func (r *authorRepository) FindAll() (authors []*models.Author, err error) {
	return authors, nil
}
