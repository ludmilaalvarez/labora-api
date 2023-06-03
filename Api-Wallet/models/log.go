package models

import "time"

type Solicitud struct {
	Id               int       `json:"id"`
	Person_id        string    `json:"national_id"`
	Date             time.Time `json:"creation_date"`
	Country          string    `json:"country"`
	Wallet_id        *int      `json:"wallet_id"`
	Status           string    `json:"status"`
	State            string    `json:"state"`
	Type_transaction string    `json:"type_transaction"`
}

type NamesFound []struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Respuesta struct {
	Check struct {
		Check_id string `json:"check_id"`
		Score    int    `json:"score"`
		Summary  struct {
			NamesFound NamesFound `json:"names_found"`
		}
	}
}
