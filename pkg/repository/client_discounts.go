package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type ClientDiscountPostgres struct {
	DB *sqlx.DB
}

func NewClientDiscountsPostgres(db *sqlx.DB) *ClientDiscountPostgres {
	return &ClientDiscountPostgres{DB: db}
}

func (s *ClientDiscountPostgres) GetAll() ([]onlinedilerv3.ClientDiscount, error) {
	var discounts []onlinedilerv3.ClientDiscount
	query := `
		SELECT cd.*, cdt.language_code, cdt.name, cdt.description
		FROM clients_discounts cd
		INNER JOIN clients_discounts_translations cdt ON cd.id = cdt.clients_discount_id
	`
	rows, err := s.DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currentDiscountID := 0
	var currentDiscount *onlinedilerv3.ClientDiscount

	for rows.Next() {
		var discount onlinedilerv3.ClientDiscount
		var translation onlinedilerv3.ClientsDiscountTrans
		err := rows.StructScan(&discount)
		if err != nil {
			return nil, err
		}

		if discount.ID != currentDiscountID {
			if currentDiscount != nil {
				discounts = append(discounts, *currentDiscount)
			}
			currentDiscount = &discount
			currentDiscountID = discount.ID
		}

		err = rows.StructScan(&translation)
		if err != nil {
			return nil, err
		}
		currentDiscount.Translations = append(currentDiscount.Translations, translation)
	}

	if currentDiscount != nil {
		discounts = append(discounts, *currentDiscount)
	}

	return discounts, nil
}

func (s *ClientDiscountPostgres) GetByID(id int) (onlinedilerv3.ClientDiscount, error) {
	var discount onlinedilerv3.ClientDiscount
	query := `
		SELECT cd.*, cdt.language_code, cdt.name, cdt.description
		FROM clients_discounts cd
		INNER JOIN clients_discounts_translations cdt ON cd.id = cdt.clients_discount_id
		WHERE cd.id = $1
	`
	rows, err := s.DB.Queryx(query, id)
	if err != nil {
		return onlinedilerv3.ClientDiscount{}, err
	}
	defer rows.Close()

	currentDiscountID := 0
	for rows.Next() {
		var translation onlinedilerv3.ClientsDiscountTrans
		err := rows.StructScan(&discount)
		if err != nil {
			return onlinedilerv3.ClientDiscount{}, err
		}

		if discount.ID != currentDiscountID {
			currentDiscountID = discount.ID
		}

		err = rows.StructScan(&translation)
		if err != nil {
			return onlinedilerv3.ClientDiscount{}, err
		}
		discount.Translations = append(discount.Translations, translation)
	}

	return discount, nil
}

func (s *ClientDiscountPostgres) Create(input onlinedilerv3.ClientDiscount) (int, error) {
	var discountID int
	query := `
		INSERT INTO clients_discounts (client_id, consignment_id, precent)
		VALUES (:client_id, :consignment_id, :percent)
		RETURNING id
	`
	namedParams := map[string]interface{}{
		"client_id":      input.ClientID,
		"consignment_id": input.ConsignmenID,
		"percent":        input.Percent,
	}
	err := s.DB.Get(&discountID, query, namedParams)
	if err != nil {
		return 0, err
	}

	for _, translation := range input.Translations {
		translationQuery := `
			INSERT INTO clients_discounts_translations (clients_discount_id, language_code, name, description)
			VALUES (:clients_discount_id, :language_code, :name, :description)
		`
		translationParams := map[string]interface{}{
			"clients_discount_id": discountID,
			"language_code":       translation.LanguageCode,
			"name":                translation.Name,
			"description":         translation.Description,
		}
		_, err := s.DB.NamedExec(translationQuery, translationParams)
		if err != nil {
			return 0, err
		}
	}

	return discountID, nil
}

func (s *ClientDiscountPostgres) Update(id int, input onlinedilerv3.ClientDiscount) error {
	query := `
		UPDATE clients_discounts
		SET client_id = :client_id, consignment_id = :consignment_id, precent = :percent
		WHERE id = :id
	`
	namedParams := map[string]interface{}{
		"id":             id,
		"client_id":      input.ClientID,
		"consignment_id": input.ConsignmenID,
		"percent":        input.Percent,
	}
	_, err := s.DB.NamedExec(query, namedParams)
	if err != nil {
		return err
	}

	for _, translation := range input.Translations {
		translationQuery := `
			UPDATE clients_discounts_translations
			SET name = :name, description = :description
			WHERE clients_discount_id = :clients_discount_id AND language_code = :language_code
		`
		translationParams := map[string]interface{}{
			"clients_discount_id": id,
			"language_code":       translation.LanguageCode,
			"name":                translation.Name,
			"description":         translation.Description,
		}
		_, err := s.DB.NamedExec(translationQuery, translationParams)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ClientDiscountPostgres) Delete(id int) error {
	query := "DELETE FROM clients_discounts WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	translationQuery := "DELETE FROM clients_discounts_translations WHERE clients_discount_id = $1"
	_, err = s.DB.Exec(translationQuery, id)
	if err != nil {
		return err
	}

	return nil
}
