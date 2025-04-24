package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"bookstore_backend/db"
	"bookstore_backend/model"
	"bookstore_backend/utils"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	// Cek Method cocok atau tidak
	if r.Method != http.MethodPost {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, r.Method, "Method not Allowed. MUST ==> POST")
		return
	}

	// Ambil data Book
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Bad request")
		return
	}

	// Eksekusi query
	query := ("INSERT INTO book (name, description, price) VALUES ($1, $2, $3)")
	_, err := db.Conn.Exec(context.Background(), query, book.Name, book.Description, book.Price)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Gagal simpan data")
		return
	}

	// Response
	utils.JSONResponse(w, http.StatusCreated, r.Method, "Data buku berhasil dibuat ✅")
}
