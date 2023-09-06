package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func GetDb(connStr string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
