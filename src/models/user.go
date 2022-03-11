package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  []byte `json:"Password"`
	IsAdmin   bool   `json:"IsAdmin"`
	Id        string `json:"_Id,omitempty"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte((password)))
}
