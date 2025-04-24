package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func InitDB(username, password, host, port, dbName, schema string) {
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s",
		username, password, host, port, dbName, schema)

	var err error
	Conn, err = pgx.Connect(context.Background(), connURL)
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}
}
