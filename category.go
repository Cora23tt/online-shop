package onlinedilerv3

type Category struct {
	ID           int                    `json:"id" db:"id"`
	Translations []CategoryTranslations `json:"translations"`
}

type CategoryTranslations struct {
	ID           int    `json:"id" db:"id"`
	CategoryID   int    `json:"category_id" db:"category_id"`
	LanguageCode string `json:"language_code" db:"language_code"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
}
