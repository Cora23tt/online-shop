package onlinedilerv3

type Order struct {
	ID           int                 `json:"id" db:"id"`
	ClientID     int                 `json:"client_id" db:"client_id"`
	OrderDate    string              `json:"order_date" db:"order_date"`
	Total        string              `json:"total" db:"total"`
	Status       string              `json:"status" db:"status"`
	Translations []OrderTranslations `json:"translations"`
}

type OrderTranslations struct {
	ID           int    `json:"id" db:"id"`
	OrderID      int    `json:"order_id" db:"order_id"`
	LanguageCode string `json:"language_code" db:"language_code"`
}

type OrderItem struct {
	ID           int     `json:"id" db:"id"`
	OrderID      int     `json:"order_id" db:"order_id"`
	ProductID    int     `json:"product_id" db:"product_id"`
	ConsignmenID int     `json:"consignment_id" db:"consignment_id"`
	Quantity     int     `json:"quantity" db:"quantity"`
	Price        float64 `json:"price" db:"price"`
}
