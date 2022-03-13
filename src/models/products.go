package models

type Product struct {
	Id          string  `json:"_Id,omitempty"`
	Title       string  `json:"Title"`
	Description string  `json:"Description"`
	Image       string  `json:"Image"`
	Price       float64 `json:"Price"`
}
