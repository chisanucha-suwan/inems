package services

type UserResponse struct {
	UserID      int    `json:"userId"`
	Username    string `json:"username"`
	FullName_TH string `json:"fullname"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(id int) (*UserResponse, error)
}
