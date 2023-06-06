package services

import (
	"Api-Wallet/models"
	"errors"

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

func WalletExists(dni string) bool {
	var count int
	query := "SELECT COUNT(*) FROM wallets WHERE person_id=$1"

	err := db.Db.QueryRow(query, dni).Scan(&count)
	if err != nil {
		log.Println(err)
	}

	return count == 1

}

func IsValidAmount(person_id string, amount float64) bool {
	var amountActual float64

	query := "SELECT  amount	FROM wallets where person_id=$1"

	err := db.Db.QueryRow(query, person_id).Scan(&amountActual)

	if err != nil {
		log.Println(err)

	}

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

func BuscarIDPersona(person_id string) float64 {
	var amount float64

	query := "SELECT  amount	FROM wallets WHERE person_id=$1 "

	err := db.Db.QueryRow(query, person_id).Scan(&amount)

	if err != nil {
		log.Println(err)

	}

	return amount

}

func (p *PostgresWallet) DeleteWallet(person_id string) error {

	err := HistoryDeleteWallet(person_id)
	if err != nil {
		return err
	}

	sqlStatement := "DELETE  FROM wallets WHERE person_id=$1"
	row, err := p.Db.Exec(sqlStatement, person_id)
	if err != nil {
		log.Println(err)
		return errors.New("No se encontro una billetera con ese documento")
	}

	count, err := row.RowsAffected()
	if err != nil {
		return errors.New("No se elimino la billetera")
	}

	if count == 0 {
		return errors.New("No se encontro una billetera con ese documento")
	}

	return nil
}

func HistoryDeleteWallet(person_id string) error {
	err := LogDelete(person_id)
	if err != nil {
		return err
	}

	err = TransactionDeleteReceiverSender(person_id)
	if err != nil {
		return err
	}

	err = TransactionUpdateTransfer(person_id)
	if err != nil {
		return err
	}

	return nil

}
