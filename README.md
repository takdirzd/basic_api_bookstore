# Aplikasi sederhana Bookstore Backend API (Go + PostgreSQL)  
  
Ini adalah aplikasi **backend sederhana** untuk toko buku (Bookstore) yang dibangun menggunakan **Go** dan **PostgreSQL**. Proyek ini mendukung berbagai operasi CRUD dasar melalui RESTful API, dan untuk frontend menggunakan **Python** dan Django sebagai framework web.  
  
---  
  
## 🚀 Fitur Utama  
  
- **Create Book**: Menambahkan data buku baru  
- **Read Books**: Menampilkan daftar semua buku  
- **Update Book**: Mengubah status atau detail buku  
- **Delete Book**: Menghapus data buku berdasarkan ID  
- **Validasi Method**: Menolak method selain yang diizinkan  
- **.env Integration**: Konfigurasi database menggunakan file `.env`  
- **Custom Response Struct**: Format response seperti `status`, `method`, dan `description`  
  
---  
  
## Struktur Direktori

```
backend/  
├── db/               # Inisialisasi koneksi database  
│   └── db.go  
├── handler/          # Handler endpoint API (CRUD)  
│   ├── create.go  
│   ├── delete.go  
│   ├── get.go  
│   └── update.go  
├── model/            # Struktur model data  
│   └── models.go  
├── utils/            # Fungsi utility umum (seperti response JSON standar)  
│   └── response.go  
├── .env              # File konfigurasi environment  
├── go.mod            # Informasi module Go dan dependency  
├── go.sum            # Checksum dependencies  
└── main.go           # Entry point aplikasi  
```
```
├── frontend/                      # Frontend Django
│   ├── bookstore_frontend/
│   │   ├── __init__.py
│   │   ├── asgi.py
│   │   ├── settings.py
│   │   ├── urls.py
│   │   ├── views.py
│   │   └── wsgi.py
│   ├── templates/
│   │   └── index.html
│   ├── db.sqlite3
│   └── manage.py
```

![image](https://github.com/user-attachments/assets/9b283c20-355e-4679-8711-26fa97a1523e)


