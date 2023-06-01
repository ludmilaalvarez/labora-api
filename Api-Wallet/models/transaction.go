package models

import "time"

type Transaction struct {
	Id                 int       `json:"id"`
	Sender_id          string    `json:"sender_id"`
	Receiver_id        string    `json:"receiver_id"`
	Amount             float64   `json:"amount"`
	Type               string    `json:"type_transaction"`
	Date               time.Time `json:"date"`
	Sender_Wallet_id   int       `json:"sender_wallet_id"`
	Receiver_Wallet_id int       `json:"receiver_wallet_id"`
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
