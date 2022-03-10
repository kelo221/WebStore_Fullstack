package models

type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  []byte `json:"Password"`
	IsAdmin   bool   `json:"IsAdmin"`
	Id        string `json:"_Id"`
}
