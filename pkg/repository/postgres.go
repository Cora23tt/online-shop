package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {

	connStr := "user=online_shop_db_user password=mIYqiES9yiOhsMvmG1ubd8lQ200BXYPQ dbname=online_shop_db host=dpg-ck7df1fsasqs73af8m1g-a sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
