package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bookstore/db"
	"bookstore/handler"
	"bookstore/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ❌")
	}

	// Ambil dari .env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	schema := os.Getenv("DB_SCHEMA")

	db.InitDB(user, pass, host, port, name, schema)

	fmt.Println("✅ Server running on http://localhost:8080")

	// URL endpoint
	http.HandleFunc("/go/books/get", utils.WithCORS(handler.GetBooksHandler))      // GET
	http.HandleFunc("/go/books/create", utils.WithCORS(handler.CreateBookHandler)) // POST
	http.HandleFunc("/go/books/update", utils.WithCORS(handler.UpdateBookHandler)) // PUT
	http.HandleFunc("/go/books/delete", utils.WithCORS(handler.DeleteBookHandler)) // DELETE

	log.Fatal(http.ListenAndServe(":8080", nil))
}
