package models

type User struct {
	UserID       int    `db:"ID"`
	Username     string `db:"Username"`
	Prefix_TH    string `db:"Prefix_TH"`
	FirstName_TH string `db:"FirstName_TH"`
	LastName_TH  string `db:"LastName_TH"`
	FullName_TH  string `db:"FullName_TH"`
}
