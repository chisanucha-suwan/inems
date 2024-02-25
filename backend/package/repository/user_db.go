package repository

import (
	"github.com/jmoiron/sqlx"
)

type usersRepositoryDB struct {
	db *sqlx.DB
}

// constructor
func NewUserRepositoryDB(db *sqlx.DB) usersRepositoryDB {
	return usersRepositoryDB{db: db}
}

func (r usersRepositoryDB) GetAll() ([]User, error) {
	users := []User{}

	query := "SELECT ID, Username, Prefix_TH, FirstName_TH, LastName_TH, FullName_TH FROM Users"
	err := r.db.Select(&users, query)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r usersRepositoryDB) GetById(id int) (*User, error) {
	user := User{}
	query := "SELECT ID, Username, Prefix_TH, FirstName_TH, LastName_TH, FullName_TH FROM Users WHERE ID = ?"
	err := r.db.Get(&user, query, id)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
