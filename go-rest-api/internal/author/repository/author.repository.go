package repository

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/leonardo-jesus/go-library-crud/go-rest-api/internal/author/models"
)

const LIMIT = 10
const ONE_MILLION = 1000000
const CHUNK_SIZE = ONE_MILLION

var END_OF_FILE_ERROR error = io.EOF

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
	csvFile, err := os.Open("./internal/author/assets/authors.csv")
	if err != nil {
		return fmt.Errorf("os.Open %w", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	for {
		chunk, isEndOfFile, err := readChunk(csvReader)
		if err != nil {
			return fmt.Errorf("error reading CSV chunk: %w", err)
		}

		err = bulkInsertOnDatabase(tx, chunk)
		if err != nil {
			return fmt.Errorf("error inserting CSV chunk: %w", err)
		}

		if isEndOfFile {
			break
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func readChunk(reader *csv.Reader) ([][]string, bool, error) {
	chunk := make([][]string, 0, CHUNK_SIZE)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == END_OF_FILE_ERROR {
				break
			}

			return nil, false, err
		}

		chunk = append(chunk, record)

		if hasDataToProccess(chunk) {
			return chunk, false, nil
		}
	}

	return chunk, true, nil
}

func hasDataToProccess(chunk [][]string) bool {
	return len(chunk) == CHUNK_SIZE
}

type CopyDataReader struct {
	chunks [][]string
	index  int
}

func (cdr *CopyDataReader) Next() bool {
	return cdr.index < len(cdr.chunks)
}

func (cdr *CopyDataReader) Values() ([]interface{}, error) {
	if cdr.index >= len(cdr.chunks) {
		return nil, END_OF_FILE_ERROR
	}

	row := cdr.chunks[cdr.index]
	cdr.index++
	values := make([]interface{}, len(row))
	for i, value := range row {
		values[i] = interface{}(value)
	}

	return values, nil
}

func (cdr *CopyDataReader) Err() error {
	return nil
}

func buildCopyData(chunks [][]string) pgx.CopyFromSource {
	return &CopyDataReader{
		chunks: chunks,
		index:  0,
	}
}

func bulkInsertOnDatabase(tx pgx.Tx, chunk [][]string) (err error) {
	_, err = tx.CopyFrom(context.Background(), pgx.Identifier{"authors"}, []string{"name"}, buildCopyData(chunk))

	return err
}
