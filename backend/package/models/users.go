package models

import "time"

// database model
type User struct {
	UserID        int        `db:"ID"`
	Username      string     `db:"Username"`
	Password      string     `db:"password"`
	PrefixTH      *string    `db:"PrefixTH"`
	FirstNameTH   string     `db:"FirstNameTH"`
	LastNameTH    string     `db:"LastNameTH"`
	FullNameTH    string     `db:"FullNameTH"`
	PrefixEN      *string    `db:"PrefixEN"`
	FirstNameEN   string     `db:"FirstNameEN"`
	LastNameEN    string     `db:"LastNameEN"`
	FullNameEN    string     `db:"FullNameEN"`
	Email         *string    `db:"Email"`
	MobilePhone   *string    `db:"MobilePhone"`
	ImgProfile    string     `db:"ImgProfile"`
	IsMember      bool       `db:"IsMember"`
	Bio           *string    `db:"Bio"`
	BioImageCSV   *string    `db:"BioImageCSV"`
	IsActive      bool       `db:"IsActive"`
	Role          int        `db:"Role"`
	PermissionCSV *string    `db:"PermissionCSV"`
	CreatedBy     string     `db:"CreatedBy"`
	CreatedDate   time.Time  `db:"CreatedDate"`
	UpdatedBy     *string    `db:"ModifiedBy"`
	UpdateDate    *time.Time `db:"ModifiedDate"`
}

// service models
type UserResponse struct {
	UserID      int        `json:"userId"`
	Username    string     `json:"username"`
	FullNameTH  string     `json:"fullname_th"`
	FullNameEN  string     `json:"fullname_en"`
	Email       *string    `json:"email"`
	MobliePhone *string    `json:"mobilephone"`
	ImgProfile  string     `json:"imgprofile"`
	IsMember    bool       `json:"ismember"`
	Role        string     `json:"role"`
	Permission  *string    `json:"permission"`
	UpdatedBy   *string    `json:"updated_by"`
	UpdatedDate *time.Time `json:"updated_date"`
}
