package onlinedilerv3

type Consignment struct {
	ID             int     `json:"id" db:"id"`
	ProductID      int     `json:"product_id" db:"product_id"`
	ExpirationDate string  `json:"expiration_date" db:"expiration_date"`
	Quantity       int     `json:"quantity" db:"quantity"`
	Price          float64 `json:"price" db:"price"`
	Description    string  `json:"description" db:"description"`
}
