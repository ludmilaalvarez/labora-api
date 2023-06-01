package services

import (
	"Api-Wallet/models"

	//"errors"
	"fmt"
	"log"
	"time"

	"Api-Wallet/db"
)

type PostgresWallet struct {
	Db db.DbConnection
}

func (p *PostgresWallet) CrearWallet(Datos *models.Datos_Solicitados) (int, error) {
	var wallet models.Wallets
	var id int
	var cols = "(person_id, date, country, state, amount, name)"
	var values = "($1, $2, $3, $4, $5, $6)"

	wallet = models.Wallets{
		Id:        id,
		Person_id: Datos.National_id,
		Date:      time.Now(),
		Country:   Datos.Country,
		State:     Datos.State,
		Amount:    0,
		Name:      Datos.Name,
	}

	var query = fmt.Sprintf("INSERT INTO wallets %s VALUES %s RETURNING id", cols, values)
	err := p.Db.QueryRow(query, wallet.Person_id, wallet.Date, wallet.Country, wallet.State, wallet.Amount, wallet.Name).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err

	}
	return id, nil
}

func (p *PostgresWallet) StatusWallet(Dni string) (models.Wallets, error) {
	var wallet models.Wallets
	query := "SELECT * FROM wallets WHERE person_id=$1"

	row, err := p.Db.Query(query, Dni)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	for row.Next() {

		err = row.Scan(&wallet.Id, &wallet.Person_id, &wallet.Date, &wallet.Country, &wallet.Amount, &wallet.Name)
		if err != nil {
			log.Println(err)
			return wallet, err
		}
	}
	return wallet, nil
}

func ComprobarExistenciaWallet(dni string) bool {
	var count int
	query := "SELECT COUNT(*) FROM wallets WHERE person_id=$1"

	err := db.Db.QueryRow(query, dni).Scan(&count)

	if err != nil {
		log.Println(err)
	}

	return count == 0

}

func VerificarMonto(person_id string, amount float64) bool {
	var amountActual float64

	query := "SELECT  amount	FROM wallets where person_id=$1"

	err := db.Db.QueryRow(query, person_id).Scan(&amountActual)

	if err != nil {
		log.Println(err)
		//return false, err
	}
	fmt.Println(amountActual >= amount)

	return (amountActual >= amount)
}

func BuscarIDWallet(person_id string) (int, string, string) {
	var (
		id      int
		country string
		state   string
	)

	query := "SELECT  id, country, state	FROM wallets where person_id=$1 "

	err := db.Db.QueryRow(query, person_id).Scan(&id, &country, &state)

	if err != nil {
		log.Println(err)

	}

	return id, country, state

}

func BuscarIDPersona(wallet_id int) (string, float64) {
	var (
		person_id string
		amount    float64
	)

	query := "SELECT  person_id, amount	FROM wallets WHERE id=$1 "

	err := db.Db.QueryRow(query, wallet_id).Scan(&person_id, &amount)

	if err != nil {
		log.Println(err)

	}

	return person_id, amount

}
