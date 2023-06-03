package models

import "time"

type Transaction struct {
	Sender_id   string    `json:"sender_id"`
	Receiver_id string    `json:"receiver_id"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type_transaction"`
	Date        time.Time `json:"date"`
}

type Movement struct {
	Type   string    `json:"type"`
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

type Transacion_respuesta struct {
	ID        string     `json:"person_id"`
	Balance   float64    `json:"balance"`
	Movements []Movement `json:"movements"`
}
