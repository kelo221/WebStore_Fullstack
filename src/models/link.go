package models

type Link struct {
	Id       string    `json:"_Id,omitempty"`
	Code     string    `json:"Code"`
	UserId   string    `json:"User_id"`
	User     User      `json:"user"`
	Products []Product `json:"Products"`
}
