package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type DiscountPostgres struct {
	DB *sqlx.DB
}

func NewDiscountPostgres(db *sqlx.DB) *DiscountPostgres {
	return &DiscountPostgres{DB: db}
}

func (s *DiscountPostgres) GetAll() ([]onlinedilerv3.Discount, error) {
	var discounts []onlinedilerv3.Discount
	query := `
		SELECT d.id, d.consignments_id, d.name, d.percent, d.description,
		       dt.language_code, dt.name as translation_name, dt.description as translation_description
		FROM discounts d
		LEFT JOIN discounts_translations dt ON d.id = dt.discount_id
	`

	rows, err := s.DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discountsMap := make(map[int]onlinedilerv3.Discount)
	for rows.Next() {
		var discount onlinedilerv3.Discount
		var translation onlinedilerv3.DiscountTranslations
		err := rows.Scan(
			&discount.ID, &discount.ConsignmentsID, &discount.Name, &discount.Percent, &discount.Description,
			&translation.LanguageCode, &translation.Name, &translation.Description,
		)
		if err != nil {
			return nil, err
		}

		if existingDiscount, ok := discountsMap[discount.ID]; ok {
			existingDiscount.Translations = append(existingDiscount.Translations, translation)
			discountsMap[discount.ID] = existingDiscount
		} else {
			discount.Translations = []onlinedilerv3.DiscountTranslations{translation}
			discountsMap[discount.ID] = discount
		}
	}

	for _, discount := range discountsMap {
		discounts = append(discounts, discount)
	}

	return discounts, nil
}

func (s *DiscountPostgres) GetByID(id int) (onlinedilerv3.Discount, error) {
	var discount onlinedilerv3.Discount
	query := `
		SELECT d.id, d.consignments_id, d.name, d.percent, d.description,
		       dt.language_code, dt.name as translation_name, dt.description as translation_description
		FROM discounts d
		LEFT JOIN discounts_translations dt ON d.id = dt.discount_id
		WHERE d.id = $1
	`

	rows, err := s.DB.Queryx(query, id)
	if err != nil {
		return onlinedilerv3.Discount{}, err
	}
	defer rows.Close()

	translations := make([]onlinedilerv3.DiscountTranslations, 0)
	for rows.Next() {
		var translation onlinedilerv3.DiscountTranslations
		err := rows.Scan(
			&discount.ID, &discount.ConsignmentsID, &discount.Name, &discount.Percent, &discount.Description,
			&translation.LanguageCode, &translation.Name, &translation.Description,
		)
		if err != nil {
			return onlinedilerv3.Discount{}, err
		}

		translations = append(translations, translation)
	}

	discount.Translations = translations
	return discount, nil
}

func (s *DiscountPostgres) Create(input onlinedilerv3.DiscountInput) (int, error) {
	// Start a transaction
	tx, err := s.DB.Beginx()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	// Insert into discounts table
	discountQuery := `
		INSERT INTO discounts (consignments_id, name, percent, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var discountID int
	err = tx.Get(&discountID, discountQuery, input.ConsignmentID, input.Name, input.Percent, input.Description)
	if err != nil {
		return 0, err
	}

	// Insert into discounts_translations table for each translation
	translationsQuery := `
		INSERT INTO discounts_translations (discount_id, language_code, name, description)
		VALUES ($1, $2, $3, $4)
	`
	for _, translation := range input.Translations {
		_, err := tx.Exec(translationsQuery, discountID, translation.LanguageCode, translation.Name, translation.Description)
		if err != nil {
			return 0, err
		}
	}

	return discountID, nil
}

func (s *DiscountPostgres) Update(id int, input onlinedilerv3.DiscountInput) error {
	// Start a transaction
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	// Update discounts table
	discountUpdateQuery := `
		UPDATE discounts
		SET consignments_id = $2, name = $3, percent = $4, description = $5
		WHERE id = $1
	`
	_, err = tx.Exec(discountUpdateQuery, id, input.ConsignmentID, input.Name, input.Percent, input.Description)
	if err != nil {
		return err
	}

	// Delete existing translations for the discount
	translationsDeleteQuery := `
		DELETE FROM discounts_translations
		WHERE discount_id = $1
	`
	_, err = tx.Exec(translationsDeleteQuery, id)
	if err != nil {
		return err
	}

	// Insert new translations for the discount
	translationsInsertQuery := `
		INSERT INTO discounts_translations (discount_id, language_code, name, description)
		VALUES ($1, $2, $3, $4)
	`
	for _, translation := range input.Translations {
		_, err := tx.Exec(translationsInsertQuery, id, translation.LanguageCode, translation.Name, translation.Description)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DiscountPostgres) Delete(id int) error {
	// Start a transaction
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	// Delete from discounts_translations table
	translationsDeleteQuery := `
		DELETE FROM discounts_translations
		WHERE discount_id = $1
	`
	_, err = tx.Exec(translationsDeleteQuery, id)
	if err != nil {
		return err
	}

	// Delete from discounts table
	discountDeleteQuery := `
		DELETE FROM discounts
		WHERE id = $1
	`
	_, err = tx.Exec(discountDeleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
