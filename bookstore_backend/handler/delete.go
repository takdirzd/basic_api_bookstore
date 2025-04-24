package handler

import (
	"context"
	"fmt"
	"net/http"

	"bookstore_backend/db"
	"bookstore_backend/utils"
)

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Cek Method cocok atau tidak
	if r.Method != http.MethodDelete {
		utils.JSONResponse(w, http.StatusMethodNotAllowed, r.Method, "Method not Allowed. MUST ==> DELETE")
		return
	}

	// Cek parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, r.Method, "Missing parameter")
		return
	}

	// Eksekusi query
	query := "DELETE FROM book WHERE id = $1"
	_, err := db.Conn.Exec(context.Background(), query, id)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, r.Method, "Gagal delete data")
		return
	}

	// Response
	utils.JSONResponse(w, http.StatusOK, r.Method, fmt.Sprintf("Data buku dengan ID %s berhasil dihapus ✅", id))
}
