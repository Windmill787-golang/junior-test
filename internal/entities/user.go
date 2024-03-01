package entities

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	PasswordHash string `json:"-"`
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
