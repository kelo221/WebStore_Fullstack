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
	Name      string `json:"Name,omitempty"`

	Address  string `json:"Address"`
	City     string `json:"City"`
	Country  string `json:"Country"`
	Zip      string `json:"Zip"`
	Complete bool   `json:"Complete"`

	Total      float64     `json:"Total,omitempty"`
	OrderItems []OrderItem `json:"OrderItem"`
}

// OrderItem A singular object that holds how many of which product has been selected.
type OrderItem struct {
	Id           string  `json:"_Id,omitempty"`
	ProductTitle string  `json:"ProductTitle"`
	Price        float64 `json:"Price"`
	Quantity     uint    `json:"Quantity"`
}

func (order *ShoppingCart) FullName() string {
	return order.FirstName + " " + order.LastName
}

func (order *ShoppingCart) CalculateTotal() float64 {
	var total float64 = 0

	for _, OrderItem := range order.OrderItems {
		total += OrderItem.Price * float64(OrderItem.Quantity)
	}

	return total
}
