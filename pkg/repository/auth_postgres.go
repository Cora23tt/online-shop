package repository

import (
	"fmt"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	DB *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{DB: db}
}

func (s *AuthPostgres) CreateUser(user onlinedilerv3.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password_hash) VALUES($1, $2, $3, $4) RETURNING id", usersTable)
	row := s.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (onlinedilerv3.User, error) {
	var user onlinedilerv3.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.DB.Get(&user, query, email, password)
	return user, err
}
