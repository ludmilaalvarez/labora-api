package models

type Item struct {
	Id           int     `json:"id"`
	CustomerName string  `json:"name"`
	OrderDate    string  `json:"date"`
	Product      string  `json:"product"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	ItemDetails  string  `json:"itemdetails"`
}
