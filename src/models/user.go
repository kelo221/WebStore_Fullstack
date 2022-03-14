package models

import (
	"golang.org/x/crypto/bcrypt"
)

type UserType int64

const (
	Regular UserType = iota
	Admin
)

type User struct {
	FirstName string `json:"FirstName,omitempty"`
	LastName  string `json:"LastName,omitempty"`
	Email     string `json:"Email,omitempty"`
	Password  []byte `json:"Password,omitempty"`
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
