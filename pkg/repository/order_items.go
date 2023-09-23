package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type OrderItemsPostgres struct {
	DB *sqlx.DB
}

func NewOrderItemsPostgres(db *sqlx.DB) *OrderItemsPostgres {
	return &OrderItemsPostgres{DB: db}
}

func (s *OrderItemsPostgres) GetItems(orderID int) ([]onlinedilerv3.OrderItem, error) {
	var items []onlinedilerv3.OrderItem

	query := `
		SELECT * FROM order_items WHERE order_id = $1
	`
	err := s.DB.Select(&items, query, orderID)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *OrderItemsPostgres) Add(orderID int, input onlinedilerv3.OrderItem) ([]onlinedilerv3.OrderItem, error) {
	query := `
		INSERT INTO order_items (order_id, product_id, consignment_id, quantity, price)
		VALUES (:order_id, :product_id, :consignment_id, :quantity, :price)
		RETURNING id
	`
	namedParams := map[string]interface{}{
		"order_id":       orderID,
		"product_id":     input.ProductID,
		"consignment_id": input.ConsignmenID,
		"quantity":       input.Quantity,
		"price":          input.Price,
	}

	var addedItemID int
	err := s.DB.Get(&addedItemID, query, namedParams)
	if err != nil {
		return nil, err
	}

	updatedItems, err := s.GetItems(orderID)
	if err != nil {
		return nil, err
	}

	return updatedItems, nil
}

func (s *OrderItemsPostgres) Update(orderID int, input onlinedilerv3.OrderItem) error {
	query := `
		UPDATE order_items
		SET product_id = :product_id, consignment_id = :consignment_id, quantity = :quantity, price = :price
		WHERE order_id = :order_id AND id = :id
	`
	namedParams := map[string]interface{}{
		"order_id":       orderID,
		"id":             input.ID,
		"product_id":     input.ProductID,
		"consignment_id": input.ConsignmenID,
		"quantity":       input.Quantity,
		"price":          input.Price,
	}

	_, err := s.DB.NamedExec(query, namedParams)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderItemsPostgres) Delete(orderID int, itemID int) error {
	query := "DELETE FROM order_items WHERE order_id = $1 AND id = $2"
	_, err := s.DB.Exec(query, orderID, itemID)
	if err != nil {
		return err
	}

	return nil
}
