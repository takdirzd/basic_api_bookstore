package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"bookstore/db"
	"bookstore/model"
	"bookstore/utils"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	//Cek Method cocok atau tidak
	if r.Method != http.MethodGet {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, r.Method, "Method Not Allowed. MUST ==> GET")
		return
	}

	// Query ke database
	rows, err := db.Conn.Query(context.Background(),
		"SELECT id, name, description, price, status, qty FROM book")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Failed Get data")
		return
	}
	defer rows.Close()

	// Masukkan hasil query ke struct
	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Description, &book.Price, &book.Status, &book.Qty); err != nil {
			utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Failed scan data")
			return
		}
		books = append(books, book)
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
