package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
)

const LIMIT = 10

type BookRepositoryInterface interface {
	FindAll(page int) (book []*models.BookDB, err error)
	FindByName(name string, page int) (book []*models.BookDB, err error)
	Create(book *models.BookAPI) (err error)
}

type bookRepository struct {
	db *pgx.Conn
}

func NewBookRepository(db *pgx.Conn) BookRepositoryInterface {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(page int) (books []*models.BookDB, err error) {
	query := `SELECT b.id, b.name, b.edition, b.publication_year, array_agg(a.name) AS authors
		FROM books b
		JOIN book_authors ba ON ba.book_id = b.id
		JOIN authors a ON a.id = ba.author_id
		GROUP BY b.id
        ORDER BY b.id
		OFFSET (($1 - 1) * $2) ROWS FETCH NEXT $2 ROWS ONLY
	`
	rows, _ := r.db.Query(context.Background(), query, page, LIMIT)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var foundBooks []*models.BookDB

	for rows.Next() {
		var book models.BookDB
		var authors []string

		err := rows.Scan(&book.Id, &book.Name, &book.Edition, &book.PublicationYear, &authors)
		if err != nil {
			log.Fatal(err)
		}

		book.Authors = authors
		foundBooks = append(foundBooks, &book)
	}

	return foundBooks, nil
}

func (r *bookRepository) FindByName(name string, page int) (books []*models.BookDB, err error) {
	query := "SELECT * FROM books WHERE name = $1 ORDER BY id OFFSET (($2 - 1) * $3) ROWS FETCH NEXT $3 ROWS ONLY"
	rows, err := r.db.Query(context.Background(), query, name, page, LIMIT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundBooks []*models.BookDB

	for rows.Next() {
		var row models.BookDB

		err := rows.Scan(&row.Id, &row.Name)
		if err != nil {
			return nil, err
		}

		foundBooks = append(foundBooks, &row)
	}

	return foundBooks, nil
}

func (r *bookRepository) Create(book *models.BookAPI) (err error) {
	sql := `
		WITH inserted_book AS (
			INSERT INTO books (name, edition, publication_year)
			VALUES ($1, $2, $3)
			RETURNING id
		)
		INSERT INTO book_authors (book_id, author_id)
		SELECT inserted_book.id, authors.id
		FROM inserted_book, authors
		WHERE authors.id = ANY($4::int[]);
	`

	_, err = r.db.Exec(context.Background(), sql, book.Name, book.Edition, book.PublicationYear, book.Authors)
	if err != nil {
		return err
	}

	return nil
}
