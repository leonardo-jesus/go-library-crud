package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/models"
)

const AuthorsCollection = "authors"

type AuthorRepositoryInterface interface {
	FindAll() (author []*models.Author, err error)
	Create() (err error)
}

type authorRepository struct {
	conn *pgx.Conn
}

func NewAuthorRepository(conn *pgx.Conn) AuthorRepositoryInterface {
	return &authorRepository{conn}
}

func (r *authorRepository) FindAll() (authors []*models.Author, err error) {
	rows, _ := r.conn.Query(context.Background(), "SELECT * FROM authors LIMIT 10")
	defer rows.Close()

	var rowSlice []models.Author
	for rows.Next() {
		var r models.Author
		err := rows.Scan(&r.Id, &r.Name)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}

	return authors, nil
}

func (r *authorRepository) Create() (err error) {
	_, err = r.conn.Exec(context.Background(), "INSERT INTO authors(name) VALUES($1)", "TESTE")

	return err
}
