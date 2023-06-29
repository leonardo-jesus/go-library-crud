package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// type connection struct {
// 	db *sql.DB
// }

func NewConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABSE_URL"))
	if err != nil {
		log.Panicln(err.Error())
	}

	return conn
}

// func (c *conn) Close() {
// 	c.session.Close()
// }
