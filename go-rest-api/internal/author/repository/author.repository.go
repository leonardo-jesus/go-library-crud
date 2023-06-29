package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/models"
)

const LIMIT = 10

type AuthorRepositoryInterface interface {
	FindAll(page int) (author []*models.Author, err error)
	FindByName(name string, page int) (author []*models.Author, err error)
	Create() (err error)
}

type authorRepository struct {
	db *pgx.Conn
}

func NewAuthorRepository(db *pgx.Conn) AuthorRepositoryInterface {
	return &authorRepository{db}
}

func (r *authorRepository) FindAll(page int) (authors []*models.Author, err error) {
	query := "SELECT * FROM authors ORDER BY id OFFSET (($1 - 1) * $2) ROWS FETCH NEXT $2 ROWS ONLY"
	rows, _ := r.db.Query(context.Background(), query, page, LIMIT)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var foundAuthors []*models.Author

	for rows.Next() {
		var row models.Author

		err := rows.Scan(&row.Id, &row.Name)
		if err != nil {
			log.Fatal(err)
		}

		foundAuthors = append(foundAuthors, &row)
	}

	return foundAuthors, nil
}

func (r *authorRepository) FindByName(name string, page int) (authors []*models.Author, err error) {
	query := "SELECT * FROM authors WHERE name = $1 ORDER BY id OFFSET (($2 - 1) * $3) ROWS FETCH NEXT $3 ROWS ONLY"
	rows, err := r.db.Query(context.Background(), query, name, page, LIMIT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundAuthors []*models.Author

	for rows.Next() {
		var row models.Author

		err := rows.Scan(&row.Id, &row.Name)
		if err != nil {
			return nil, err
		}

		foundAuthors = append(foundAuthors, &row)
	}

	return foundAuthors, nil
}

func (r *authorRepository) Create() (err error) {
	_, err = r.db.Exec(context.Background(), "INSERT INTO authors(name) VALUES($1)", "TESTE")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
