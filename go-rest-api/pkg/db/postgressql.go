package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicln(err.Error())
	}

	return conn
}
