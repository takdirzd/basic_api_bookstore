package handler

import (
	"context"
	"fmt"
	"net/http"

	"bookstore_backend/db"
	"bookstore_backend/utils"
)

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	// Cek Method cocok atau tidak
	if r.Method != http.MethodPut {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, r.Method, "Method not Allowed. MUST ==> PUT")
		return
	}

	// Ambil parameter
	id := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")
	if id == "" || status == "" {
		utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Missing parameter id or status")
		return
	}

	// Eksekusi query
	query := "UPDATE book SET status = $1 WHERE id = $2"
	_, err := db.Conn.Exec(context.Background(), query, status, id)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Gagal update data")
		return
	}

	// Response
	utils.JSONResponse(w, http.StatusOK, r.Method, fmt.Sprintf("Data buku dengan ID %s berhasil diupdate ✅", id))
}
