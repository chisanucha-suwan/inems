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

	query := `
		SELECT 
			ID,
			Username,
			FullNameTH,
			FullNameEN,
			Email,
			MobilePhone,
			ImgProfile,
			IsMember,
			Role,
			PermissionCSV,
			CASE
				WHEN ModifiedBy IS NULL THEN CreatedBy
				ELSE ModifiedBy
			END AS ModifiedBy,
			CASE 
				WHEN ModifiedDate IS NULL THEN CreatedDate
				ELSE ModifiedDate
			END AS ModifiedDate
		FROM dbo.Users
		WHERE IsActive = 1
		ORDER BY ID;
	`
	err := r.DB.Select(&users, query)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetById(id int) (*models.User, error) {
	user := models.User{}
	query := "SELECT ID, Username, FirstNameTH, LastNameTH, FullNameTH FROM Users WHERE ID = ?"
	err := r.DB.Get(&user, query, id)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
