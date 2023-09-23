package onlinedilerv3

type Discount struct {
	ID             int                    `json:"id" db:"id"`
	ConsignmentsID int                    `json:"consignments_id" db:"consignments_id"`
	Name           string                 `json:"name" db:"name"`
	Percent        float64                `json:"percent" db:"percent"`
	Description    string                 `json:"description" db:"description"`
	Translations   []DiscountTranslations `json:"translations" db:"translations"`
}

type DiscountTranslations struct {
	ID int `json:"id" db:"id"`
	// DiscountID   int    `json:"discount_id" db:"discount_id"`
	LanguageCode string `json:"language_code" db:"language_code"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
}

type DiscountInput struct {
	ConsignmentID int                         `json:"consignment_id" db:"consignments_id"`
	Name          string                      `json:"name" db:"name"`
	Percent       float64                     `json:"percent" db:"percent"`
	Description   string                      `json:"description" db:"description"`
	Translations  []DiscountTranslationsInput `json:"translations" db:"translations"`
}

type DiscountTranslationsInput struct {
	LanguageCode string `json:"language_code" db:"language_code"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
}

type ClientDiscount struct {
	ID           int                    `json:"id" db:"id"`
	ClientID     int                    `json:"client_id" db:"client_id"`
	ConsignmenID int                    `json:"consignment_id" db:"consignment_id"`
	Percent      float64                `json:"percent" db:"percent"`
	Translations []ClientsDiscountTrans `json:"translations" db:"translations"`
}

type ClientsDiscountTrans struct {
	LanguageCode string `json:"language_code" db:"language_code"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
}
