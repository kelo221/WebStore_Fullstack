package models

// ShoppingCart Contains information about the person AND the items ordered.
type ShoppingCart struct {
	Id            string `json:"_Id,omitempty"`
	TransactionId string `json:"TransactionId"`
	UserId        string `json:"UserId"`
	Code          string `json:"Code"`

	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`

	Address  string `json:"Address"`
	City     string `json:"City"`
	Country  string `json:"Country"`
	Zip      string `json:"Zip"`
	Complete bool   `json:"Complete"`

	OrderItem []OrderItem `json:"OrderItem"`
}

// OrderItem A singular object that holds how many of which product has been selected.
type OrderItem struct {
	Id           string  `json:"_Id,omitempty"`
	ProductTitle string  `json:"ProductTitle"`
	Price        float64 `json:"Price"`
	Quantity     uint    `json:"Quantity"`
}
