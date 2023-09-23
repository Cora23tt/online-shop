package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type ConsignmentPostgres struct {
	DB *sqlx.DB
}

func NewConsignmentPostgres(db *sqlx.DB) *ConsignmentPostgres {
	return &ConsignmentPostgres{DB: db}
}

func (s *ConsignmentPostgres) GetAll() ([]onlinedilerv3.Consignment, error) {
	var consignments []onlinedilerv3.Consignment
	query := `
		SELECT id, product_id, expiration_date, quantity, price, description
		FROM consignments
	`

	if err := s.DB.Select(&consignments, query); err != nil {
		return nil, err
	}

	return consignments, nil
}

func (s *ConsignmentPostgres) GetByID(id int) (onlinedilerv3.Consignment, error) {
	var consignment onlinedilerv3.Consignment
	query := `
		SELECT id, product_id, expiration_date, quantity, price, description
		FROM consignments
		WHERE id = $1
	`

	if err := s.DB.Get(&consignment, query, id); err != nil {
		return onlinedilerv3.Consignment{}, err
	}

	return consignment, nil
}

func (s *ConsignmentPostgres) Create(input onlinedilerv3.Consignment) (int, error) {
	query := `
		INSERT INTO consignments (product_id, expiration_date, quantity, price, description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var insertedID int
	err := s.DB.Get(&insertedID, query, input.ProductID, input.ExpirationDate, input.Quantity, input.Price, input.Description)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (s *ConsignmentPostgres) Update(id int, input onlinedilerv3.Consignment) error {
	query := `
		UPDATE consignments
		SET product_id = $2, expiration_date = $3, quantity = $4, price = $5, description = $6
		WHERE id = $1
	`

	_, err := s.DB.Exec(query, id, input.ProductID, input.ExpirationDate, input.Quantity, input.Price, input.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConsignmentPostgres) Delete(id int) error {
	query := `
		DELETE FROM consignments
		WHERE id = $1
	`

	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
