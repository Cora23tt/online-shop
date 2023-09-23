package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	DB *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{DB: db}
}

func (s *UserPostgres) Search(name string) ([]onlinedilerv3.User, error) {
	query := `
		SELECT id, first_name, last_name, email, password_hash, phone_number, latitude, longtitude
		FROM users
		WHERE first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%'
	`

	var users []onlinedilerv3.User
	if err := s.DB.Select(&users, query, name); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserPostgres) GetAll() ([]onlinedilerv3.User, error) {
	query := `
		SELECT id, first_name, last_name, email, password_hash, phone_number, latitude, longtitude
		FROM users
	`

	var users []onlinedilerv3.User
	if err := s.DB.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserPostgres) GetByID(id int) (onlinedilerv3.User, error) {
	var user onlinedilerv3.User
	query := `
		SELECT id, first_name, last_name, email, password_hash, phone_number, latitude, longtitude
		FROM users
		WHERE id = $1
	`

	if err := s.DB.Get(&user, query, id); err != nil {
		return onlinedilerv3.User{}, err
	}

	return user, nil
}

func (s *UserPostgres) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserPostgres) Update(id int, user onlinedilerv3.User) error {
	query := `
		UPDATE users
		SET first_name = $2, last_name = $3, email = $4, password_hash = $5,
		    phone_number = $6, latitude = $7, longtitude = $8
		WHERE id = $1
	`

	_, err := s.DB.Exec(query, id, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.Latitude, user.Longtitude)
	if err != nil {
		return err
	}

	return nil
}
