package model

//Struktur JSON list buku
type Book struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
	Qty         int     `json:"qty"`
}

//Struktur JSON tiap response
type Response struct {
	Status      int    `json:"status"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type CartItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
