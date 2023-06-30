package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/book/models"
)

const LIMIT = 10

type BookRepositoryInterface interface {
	FindAll(page int) (book []*models.Book, err error)
	FindByName(name string, page int) (book []*models.Book, err error)
	FindById(id int) (book *models.UpdateBookSchema, err error)
	Create(book *models.CreateBookSchema) (err error)
	Update(book *models.UpdateBookSchema) (err error)
	Delete(id int) (err error)
}

type bookRepository struct {
	db *pgx.Conn
}

func NewBookRepository(db *pgx.Conn) BookRepositoryInterface {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(page int) (books []*models.Book, err error) {
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

	var foundBooks []*models.Book

	for rows.Next() {
		var book models.Book
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

func (r *bookRepository) FindByName(name string, page int) (books []*models.Book, err error) {
	query := "SELECT * FROM books WHERE name = $1 ORDER BY id OFFSET (($2 - 1) * $3) ROWS FETCH NEXT $3 ROWS ONLY"
	rows, err := r.db.Query(context.Background(), query, name, page, LIMIT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundBooks []*models.Book

	for rows.Next() {
		var row models.Book

		err := rows.Scan(&row.Id, &row.Name)
		if err != nil {
			return nil, err
		}

		foundBooks = append(foundBooks, &row)
	}

	return foundBooks, nil
}

func (r *bookRepository) FindById(id int) (books *models.UpdateBookSchema, err error) {
	query := `
		SELECT b.id, b.name, b.edition, b.publication_year, array_agg(a.id) AS authors
		FROM books b
		JOIN book_authors ba ON ba.book_id = b.id
		JOIN authors a ON a.id = ba.author_id
		WHERE b.id = $1
		GROUP BY b.id
		ORDER BY b.id
		LIMIT 1
	`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundBook *models.UpdateBookSchema

	for rows.Next() {
		var book models.UpdateBookSchema
		var authors []int

		err := rows.Scan(&book.Id, &book.Name, &book.Edition, &book.PublicationYear, &authors)
		if err != nil {
			log.Fatal(err)
		}

		book.Authors = &authors
		foundBook = &book
	}

	return foundBook, nil
}

func (r *bookRepository) Create(book *models.CreateBookSchema) (err error) {
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

func (r *bookRepository) Update(book *models.UpdateBookSchema) (err error) {
	updateSQL := `
		UPDATE books
		SET name = $1,
			edition = $2,
			publication_year = $3
		WHERE id = $4
	`

	_, err = r.db.Exec(context.Background(), updateSQL, *book.Name, *book.Edition, *book.PublicationYear, book.Id)
	if err != nil {
		panic(err)
	}

	r.DeleteBookFromBookAuthors(book.Id)
	r.InsertBookInBookAuthors(book.Id, book.Authors)

	return nil
}

func (r *bookRepository) DeleteBookFromBookAuthors(id int) (err error) {
	deleteSQL := `
		DELETE FROM book_authors WHERE book_id = $1
	`

	_, err = r.db.Exec(context.Background(), deleteSQL, id)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *bookRepository) InsertBookInBookAuthors(bookId int, authorsId *[]int) (err error) {
	insertSQL := `
		INSERT INTO book_authors (book_id, author_id)
		SELECT $1, authors.id
		FROM authors
		WHERE authors.id = ANY($2::int[])
	`

	_, err = r.db.Exec(context.Background(), insertSQL, bookId, *authorsId)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *bookRepository) Delete(id int) (err error) {
	deleteSQL := `
		DELETE FROM books WHERE id = $1
	`

	_, err = r.db.Exec(context.Background(), deleteSQL, id)
	if err != nil {
		panic(err)
	}

	return nil
}
