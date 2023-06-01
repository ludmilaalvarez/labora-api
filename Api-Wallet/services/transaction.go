package services

import (
	"Api-Wallet/db"
	"Api-Wallet/models"
	"errors"
	"fmt"
	"log"
	"time"
	//"golang.org/x/text/cases"
)

func CreateTransaction(newTransaccion models.Transaction) error {

	var status string
	if !VerificarDocumentosTransaction(newTransaccion) {
		err := RealizarTransaction(newTransaccion)
		if err != nil {
			status = "Rechazado"
			RegistrarTransaccion(status, newTransaccion)
			return err

		}
		status = "Completado"
		err = RegistrarTransaccion(status, newTransaccion)
		if err != nil {
			return err
		}
		err = AlmacenarTransaccion(newTransaccion)
		if err != nil {
			return err
		}
	}
	return nil

}

func RealizarTransaction(newTransaccion models.Transaction) error {
	tipo_transaccion := newTransaccion.Type
	switch tipo_transaccion {
	case "deposit":

		err := Transaccion(newTransaccion.Sender_id, newTransaccion.Amount)
		if err != nil {
			log.Println(err)
			return err
		}
	case "withdrawal":

		if VerificarMonto(newTransaccion.Sender_id, newTransaccion.Amount) {
			err := Transaccion(newTransaccion.Sender_id, -(newTransaccion.Amount))
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			return errors.New("Monto insuficiente")
		}
	case "transfer":

		if VerificarMonto(newTransaccion.Sender_id, newTransaccion.Amount) {

			err := Transaccion(newTransaccion.Receiver_id, newTransaccion.Amount)
			if err != nil {
				log.Println(err)
				return err
			}
			monto := -newTransaccion.Amount
			fmt.Println(monto)
			err = Transaccion(newTransaccion.Sender_id, -(newTransaccion.Amount))
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			return errors.New("Monto insuficiente")
		}
	}
	return nil

}

func Transaccion(person_id string, amount float64) error {
	updateStatement := "UPDATE wallets SET  amount=amount+$1 WHERE person_id=$2 "
	_, err := db.Db.Exec(updateStatement, amount, person_id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

/* func Withdrawal(person_id string, amount float64)error{
	updateStatement:="UPDATE wallets SET  amount=amount-$1 WHERE person_id=$2 "
	_, err:= db.Db.Exec(updateStatement,amount, person_id)

	if err != nil{
		log.Println(err)
		return err
	}
	return nil

} */

func VerificarDocumentosTransaction(newTransaccion models.Transaction) bool {
	comprobar := ComprobarExistenciaWallet(newTransaccion.Receiver_id) && ComprobarExistenciaWallet(newTransaccion.Sender_id)

	return comprobar
}

func AlmacenarTransaccion(newTransaccion models.Transaction) error {
	switch newTransaccion.Type {
	case "transfer":
		err := AlmacenarTransferencia(newTransaccion)
		if err != nil {
			log.Println(err)
			return err
		} else {
			return nil
		}
	case "withdrawal":
		err := AlmacenarExtraccion(newTransaccion)
		return err
	case "deposit":
		err := AlmacenarDeposito(newTransaccion)
		return err
	}

	return nil
}

func AlmacenarTransferencia(newTransaccion models.Transaction) error {
	senderWallet_id, _, _ := BuscarIDWallet(newTransaccion.Sender_id)
	receiverWallet_id, _, _ := BuscarIDWallet(newTransaccion.Receiver_id)

	insertStatement := `INSERT INTO transaction (sender_id, receiver_id, amount, type, date, sender_wallet_id, receiver_wallet_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Db.Exec(insertStatement, newTransaccion.Sender_id, newTransaccion.Receiver_id, newTransaccion.Amount, newTransaccion.Type, time.Now(), senderWallet_id, receiverWallet_id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func AlmacenarExtraccion(newTransaccion models.Transaction) error {
	insertStatement := `INSERT INTO transaction (receiver_id, amount, type, date, sender_wallet_id)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Db.Exec(insertStatement, newTransaccion.Receiver_id, newTransaccion.Amount, newTransaccion.Type, time.Now(), newTransaccion.Sender_Wallet_id)

	if err != nil {
		log.Println(err)
	}
	return err
}

func AlmacenarDeposito(newTransaccion models.Transaction) error {
	insertStatement := `INSERT INTO transaction (sender_id, amount, type, date, receiver_wallet_id)
			VALUES ($1, $2, $3, $4, $5)`

	fmt.Println(newTransaccion)

	_, err := db.Db.Exec(insertStatement, newTransaccion.Sender_id, newTransaccion.Amount, newTransaccion.Type, time.Now(), newTransaccion.Receiver_Wallet_id)

	if err != nil {
		log.Println(err)
	}
	return err
}

func HistorialTransacciones(wallet_id int) (models.Transacion_respuesta, error) {
	var transactionsDetails models.Transacion_respuesta

	selectStatement := `SELECT amount, type, date FROM transaction WHERE sender_wallet_id = $1 OR receiver_wallet_id = $1`

	person_id, amount := BuscarIDPersona(wallet_id)

	transactionsDetails.ID = person_id
	transactionsDetails.Balance = amount

	row, err := db.Db.Query(selectStatement, wallet_id)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	var movements []models.Movement

	for row.Next() {
		var movement models.Movement
		err = row.Scan(&movement.Amount, &movement.Type, &movement.Date)
		if err != nil {
			log.Println(err)
			return models.Transacion_respuesta{}, err
		}
		movements = append(movements, movement)
	}
	transactionsDetails.Movements = movements

	return transactionsDetails, nil
}
