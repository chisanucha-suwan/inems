package repository

import (
	"github.com/doug-martin/goqu/v9"
)

const (
	Users = "Users"
)

type User struct {
	UserID       int     `db:"ID"`
	Username     string  `db:"Username"`
	Prefix_TH    *string `db:"Prefix_TH"`
	FirstName_TH string  `db:"FirstName_TH"`
	LastName_TH  string  `db:"LastName_TH"`
	FullName_TH  string  `db:"FullName_TH"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(id int) (*User, error)
}

type SqlRepository struct {
	DB *goqu.Database
}

func (r *SqlRepository) GetAllUser() (string, error) {
	sql, _, err := r.DB.
		From("Users").
		Select("role").
		ToSQL()

	if err != nil {
		return "", err
	}

	return sql, nil
}
