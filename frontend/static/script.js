async function getBooks() {
    const response = await fetch("http://localhost:8080/go/books/get");
    const data = await response.json();
    const tableBody = document.getElementById("book-list");
    tableBody.innerHTML = "";
    data.forEach(book => {
      const row = document.createElement("tr");
      row.innerHTML = `
        <td>${book.id}</td>
        <td>${book.name}</td>
        <td>${book.description}</td>
        <td>${book.price}</td>
        <td>${book.status}</td>
        <td>${book.qty}</td>
      `;
      tableBody.appendChild(row);
    });
  }
  
  async function addBook() {
    const name = document.getElementById("name").value;
    const description = document.getElementById("description").value;
    const price = parseFloat(document.getElementById("price").value);
  
    const bookData = { name, description, price };
  
    try {
      const response = await fetch("http://localhost:8080/go/books/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(bookData)
      });
  
      if (response.ok) {
        alert("Buku berhasil ditambahkan!");
        getBooks();
      } else {
        const error = await response.json();
        alert("Gagal menambahkan buku: " + (error.message || "Unknown error"));
      }
    } catch (error) {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat menambah buku.");
    }
  }
  
  async function updateBook() {
    const id = document.getElementById("update-id").value;
    const status = document.getElementById("update-status").value;
  
    const bookData = { status };
  
    try {
      const response = await fetch(`http://localhost:8080/go/books/update?status=${status}&id=${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(bookData)
      });
  
      if (response.ok) {
        alert("Status buku berhasil diperbarui!");
        getBooks();
      } else {
        alert("Gagal memperbarui status buku.");
      }
    } catch (error) {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat update buku.");
    }
  }
  
  async function deleteBook() {
    const id = document.getElementById("delete-id").value;
  
    const response = await fetch(`http://localhost:8080/go/books/delete?id=${id}`, {
      method: "DELETE"
    });
  
    if (response.ok) {
      alert("Buku berhasil dihapus.");
      getBooks();
    } else {
      alert("Gagal menghapus buku.");
    }
  }
  
  getBooks();
  