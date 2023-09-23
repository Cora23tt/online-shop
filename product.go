package onlinedilerv3

import "github.com/lib/pq"

type Product struct {
	ID           int                  `json:"id" db:"id"`
	Rating       float32              `json:"rating" db:"rating"`
	CategoryID   int                  `json:"category_id" db:"category_id"`
	Translations []ProductTranslation `json:"translations"`
	ImageURLs    []string             `json:"image_urls" db:"image_urls"`
}

type ProductTranslation struct {
	LanguageCode string `json:"language_code" db:"language_code"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
}

type ProductComplex struct {
	Product
	ProductTranslation
}

type ProductComplect struct {
	Product
	Translations []ProductTranslation `json:"translations"`
}

type ProductTranslationConsignment struct {
	ID                     int            `json:"id" db:"id"`
	Rating                 float32        `json:"rating" db:"rating"`
	CategoryID             int            `json:"category_id" db:"category_id"`
	LanguageCode           string         `json:"language_code" db:"language_code"`
	Name                   string         `json:"name" db:"name"`
	Description            string         `json:"description" db:"description"`
	ExpirationDate         string         `json:"expiration_date" db:"expiration_date"`
	Quantity               int            `json:"quantity" db:"quantity"`
	Price                  float64        `json:"price" db:"price"`
	ConsignmentDescription string         `json:"consignment_description" db:"consignment_description"`
	ImageURLs              pq.StringArray `json:"image_urls" db:"image_urls"`
}
