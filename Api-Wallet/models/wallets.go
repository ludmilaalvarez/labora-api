package models

import "time"

type Wallets struct {
	Id        int       `json:"id"`
	Person_id string    `json:"national_id"`
	Date      time.Time `json:"creation_date"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	Amount    float64   `json:"balance"`
	Name      string    `json:"name"`
}
