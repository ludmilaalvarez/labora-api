package models

import "time"

type Wallets struct {
	Id        int       `json:"id"`
	Person_id string    `json:"national_id"`
	Date      time.Time `json:"creation_date"`
	Country   string    `json:"country"`
	Amount    float64   `json:"amount"`
	Name      string    `json:"name"`
}
