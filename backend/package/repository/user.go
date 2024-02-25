package repository

import (
	"inems/package/models"

	"github.com/jmoiron/sqlx"
)

const (
	Users = "Users"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id int) (*models.User, error)
}

type repository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetAll() ([]models.User, error) {
	users := []models.User{}

	query := "SELECT ID, Username, FirstName_TH, LastName_TH, FullName_TH FROM Users"
	err := r.DB.Select(&users, query)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetById(id int) (*models.User, error) {
	user := models.User{}
	query := "SELECT ID, Username, FirstName_TH, LastName_TH, FullName_TH FROM Users WHERE ID = ?"
	err := r.DB.Get(&user, query, id)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
