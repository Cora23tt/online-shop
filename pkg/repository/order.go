package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	DB *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{DB: db}
}

func (s *OrderPostgres) GetAll() ([]onlinedilerv3.Order, error) {
	var orders []onlinedilerv3.Order

	query := `
		SELECT o.*, ot.language_code, ot.status
		FROM orders o
		INNER JOIN orders_translations ot ON o.id = ot.order_id
	`
	rows, err := s.DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currentOrderID := 0
	var currentOrder *onlinedilerv3.Order

	for rows.Next() {
		var translation onlinedilerv3.OrderTranslations
		var order onlinedilerv3.Order
		err := rows.StructScan(&order)
		if err != nil {
			return nil, err
		}

		if order.ID != currentOrderID {
			if currentOrder != nil {
				orders = append(orders, *currentOrder)
			}
			currentOrder = &order
			currentOrderID = order.ID
		}

		err = rows.StructScan(&translation)
		if err != nil {
			return nil, err
		}
		currentOrder.Translations = append(currentOrder.Translations, translation)
	}

	if currentOrder != nil {
		orders = append(orders, *currentOrder)
	}

	return orders, nil
}

func (s *OrderPostgres) GetByID(id int) (onlinedilerv3.Order, error) {
	var order onlinedilerv3.Order

	query := `
		SELECT o.*, ot.language_code, ot.status
		FROM orders o
		INNER JOIN orders_translations ot ON o.id = ot.order_id
		WHERE o.id = $1
	`
	rows, err := s.DB.Queryx(query, id)
	if err != nil {
		return onlinedilerv3.Order{}, err
	}
	defer rows.Close()

	currentOrderID := 0
	var currentOrder *onlinedilerv3.Order

	for rows.Next() {
		var translation onlinedilerv3.OrderTranslations
		var order onlinedilerv3.Order
		err := rows.StructScan(&order)
		if err != nil {
			return onlinedilerv3.Order{}, err
		}

		if order.ID != currentOrderID {
			if currentOrder != nil {
				order.Translations = append(order.Translations, translation)
				break
			}
			currentOrder = &order
			currentOrderID = order.ID
		}

		err = rows.StructScan(&translation)
		if err != nil {
			return onlinedilerv3.Order{}, err
		}
		currentOrder.Translations = append(currentOrder.Translations, translation)
	}

	return order, nil
}

func (s *OrderPostgres) Create(input onlinedilerv3.Order) (int, error) {
	var orderID int

	orderInsertQuery := `
		INSERT INTO orders (client_id, order_date, total, status)
		VALUES (:client_id, :order_date, :total, :status)
		RETURNING id
	`
	orderParams := map[string]interface{}{
		"client_id":  input.ClientID,
		"order_date": input.OrderDate,
		"total":      input.Total,
		"status":     input.Status,
	}
	err := s.DB.Get(&orderID, orderInsertQuery, orderParams)
	if err != nil {
		return 0, err
	}

	for _, translation := range input.Translations {
		translationInsertQuery := `
			INSERT INTO orders_translations (order_id, language_code)
			VALUES (:order_id, :language_code)
		`
		translationParams := map[string]interface{}{
			"order_id":      orderID,
			"language_code": translation.LanguageCode,
		}
		_, err := s.DB.NamedExec(translationInsertQuery, translationParams)
		if err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func (s *OrderPostgres) Update(id int, input onlinedilerv3.Order) error {
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updateOrderQuery := `
		UPDATE orders
		SET client_id = :client_id, order_date = :order_date, total = :total, status = :status
		WHERE id = :id
	`
	_, err = tx.NamedExec(updateOrderQuery, map[string]interface{}{
		"id":         id,
		"client_id":  input.ClientID,
		"order_date": input.OrderDate,
		"total":      input.Total,
		"status":     input.Status,
	})
	if err != nil {
		return err
	}

	deleteTranslationsQuery := `
		DELETE FROM orders_translations
		WHERE order_id = :order_id
	`
	_, err = tx.NamedExec(deleteTranslationsQuery, map[string]interface{}{
		"order_id": id,
	})
	if err != nil {
		return err
	}

	insertTranslationsQuery := `
		INSERT INTO orders_translations (order_id, language_code)
		VALUES (:order_id, :language_code)
	`
	for _, translation := range input.Translations {
		_, err = tx.NamedExec(insertTranslationsQuery, map[string]interface{}{
			"order_id":      id,
			"language_code": translation.LanguageCode,
		})
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *OrderPostgres) Delete(id int) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	translationQuery := "DELETE FROM orders_translations WHERE order_id = $1"
	_, err = s.DB.Exec(translationQuery, id)
	if err != nil {
		return err
	}

	return nil
}
