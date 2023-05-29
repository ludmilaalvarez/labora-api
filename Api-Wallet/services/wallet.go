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

func (p *PostgresWallet) CrearWallet(Datos *models.Datos_Solicitados) int {
	var wallet models.Wallets
	var id int
	var cols = "(person_id, date, country, amount, name)"
	var values = "($1, $2, $3, $4, $5)"

	wallet = models.Wallets{
		Id:        id,
		Person_id: Datos.National_id,
		Date:      time.Now(),
		Country:   Datos.Country,
		Amount:    0,
		Name:      Datos.Name,
	}

	var query = fmt.Sprintf("INSERT INTO wallets %s VALUES %s RETURNING id", cols, values)
	err := p.Db.QueryRow(query, wallet.Person_id, wallet.Date, wallet.Country, wallet.Amount, wallet.Name).Scan(&id)
	if err != nil {
		log.Println(err)

	}

	/* defer func() {
		if err != nil {
			Tx.Rollback()
		} else {
			err = Tx.Commit()
		}
	}() */

	return id
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
			continue
		}
	}
	return wallet, nil
}

func ComprobarExistencia(dni string) int {
	var count int
	query := "SELECT COUNT(*) FROM wallets WHERE person_id=$1"

	err := db.Db.QueryRow(query, dni).Scan(&count)

	if err != nil {
		log.Println(err)
	}

	return count

}
