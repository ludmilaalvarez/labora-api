package services

import (
	"Api-Wallet/db"
	"Api-Wallet/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func CreateTransaction(newTransaction models.Transaction) error {

	var status string

	tx, err := db.Db.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if !IsDocumentTransactionValid(newTransaction) {
		return errors.New("Documentos invalidos")
	}

	err = RealizarTransaction(newTransaction, tx)
	if err != nil {
		status = "Rechazado"
		RecordTransaction(status, newTransaction, tx)

		return err

	}

	status = "Completado"

	err = RecordTransaction(status, newTransaction, tx)
	if err != nil {
		return err
	}

	err = AlmacenarTransaccion(newTransaction, tx)
	if err != nil {

		return err
	}

	return nil
}

/*
	 func RealizarTransaction(newTransaction models.Transaction) error {
		tipo_transaccion := newTransaction.Type

		switch tipo_transaccion {
		case "deposit":

			err := IncreseAccountAmount(newTransaction.Sender_id, newTransaction.Amount)
			if err != nil {
				log.Println(err)
				return err
			}
		case "withdrawal":

			if !IsValidAmount(newTransaction.Receiver_id, newTransaction.Amount) {
				return errors.New("Monto insuficiente")
			}

			err := ReduceAccountAmount(newTransaction.Receiver_id, newTransaction.Amount)
			if err != nil {
				log.Println(err)
				return err
			}

		case "transfer":

			if !IsValidAmount(newTransaction.Sender_id, newTransaction.Amount) {
				return errors.New("Monto insuficiente")
			}

			err := ReduceAccountAmount(newTransaction.Sender_id, newTransaction.Amount)
			if err != nil {
				log.Println(err)
				return err
			}

			err = IncreseAccountAmount(newTransaction.Receiver_id, newTransaction.Amount)
			if err != nil {
				log.Println(err)
				return err
			}

		}
		return nil

}
*/
func RealizarTransaction(newTransaction models.Transaction, tx *sql.Tx) error {
	tipo_transaccion := newTransaction.Type

	acciones := map[string]func() error{
		"deposit": func() error {
			err := IncreseAccountAmount(newTransaction.Sender_id, newTransaction.Amount, tx)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
		"withdrawal": func() error {
			if !IsValidAmount(newTransaction.Receiver_id, newTransaction.Amount) {
				return errors.New("Monto insuficiente")
			}

			err := ReduceAccountAmount(newTransaction.Receiver_id, newTransaction.Amount, tx)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
		"transfer": func() error {
			if !IsValidAmount(newTransaction.Sender_id, newTransaction.Amount) {
				return errors.New("Monto insuficiente")
			}

			err := ReduceAccountAmount(newTransaction.Sender_id, newTransaction.Amount, tx)
			if err != nil {
				log.Println(err)
				return err
			}

			err = IncreseAccountAmount(newTransaction.Receiver_id, newTransaction.Amount, tx)
			if err != nil {
				log.Println(err)
				return err
			}

			return nil
		},
	}

	if accion, ok := acciones[tipo_transaccion]; ok {
		return accion()
	}

	return errors.New("Tipo de transacción inválido")
}

func IncreseAccountAmount(receiver_id string, amount float64, tx *sql.Tx) error {

	return Transaccion(receiver_id, amount, tx)
}

func ReduceAccountAmount(sender_id string, amount float64, tx *sql.Tx) error {

	return Transaccion(sender_id, -amount, tx)
}

func Transaccion(person_id string, amount float64, tx *sql.Tx) error {
	updateStatement := "UPDATE wallets SET  amount=amount+$1 WHERE person_id=$2 "
	_, err := tx.Exec(updateStatement, amount, person_id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func IsDocumentTransactionValid(newTransaccion models.Transaction) bool { //Intentar arreglar esto
	return WalletExists(newTransaccion.Receiver_id) && WalletExists(newTransaccion.Sender_id)
}

func AlmacenarTransaccion(newTransaccion models.Transaction, tx *sql.Tx) error {

	switch newTransaccion.Type {
	case "transfer":
		err := AlmacenarTransferencia(newTransaccion, tx)
		if err != nil {
			log.Println(err)
			return err
		} else {
			return nil
		}
	case "withdrawal":
		err := AlmacenarExtraccion(newTransaccion, tx)
		return err
	case "deposit":
		err := AlmacenarDeposito(newTransaccion, tx)
		return err
	}

	return nil
}

func AlmacenarTransferencia(newTransaccion models.Transaction, tx *sql.Tx) error {

	insertStatement := `INSERT INTO transaction (sender_id, receiver_id, amount, type, date)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Db.Exec(insertStatement, newTransaccion.Sender_id, newTransaccion.Receiver_id, newTransaccion.Amount, newTransaccion.Type, time.Now())
	if err != nil {
		log.Println(err)
	}
	return nil
}

func AlmacenarExtraccion(newTransaccion models.Transaction, tx *sql.Tx) error {
	insertStatement := `INSERT INTO transaction (receiver_id, amount, type, date)
	VALUES ($1, $2, $3, $4)`

	_, err := db.Db.Exec(insertStatement, newTransaccion.Receiver_id, newTransaccion.Amount, newTransaccion.Type, time.Now())

	if err != nil {
		log.Println(err)
	}
	return err
}

func AlmacenarDeposito(newTransaccion models.Transaction, tx *sql.Tx) error {
	insertStatement := `INSERT INTO transaction (sender_id, amount, type, date)
			VALUES ($1, $2, $3, $4)`

	fmt.Println(newTransaccion)

	_, err := tx.Exec(insertStatement, newTransaccion.Sender_id, newTransaccion.Amount, newTransaccion.Type, time.Now())

	if err != nil {
		log.Println(err)
	}
	return err
}

func HistorialTransacciones(person_id string) (models.Transacion_respuesta, error) {
	var transactionsDetails models.Transacion_respuesta

	//selectStatement := `SELECT amount, type, date FROM transaction WHERE sender_id = $1 OR receiver_id = $1`

	selectStatement := ` SELECT * FROM
	(SELECT 'receiver_id' AS role_string, receiver_id AS wallet_id , date , amount FROM transaction
	UNION ALL
	SELECT 'sender_id'AS role_string, sender_id AS wallet_id, date, amount FROM transaction
	ORDER BY date) as trans WHERE wallet_id = $1`

	amount := BuscarIDPersona(person_id)

	transactionsDetails.Balance = amount

	row, err := db.Db.Query(selectStatement, person_id)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	var movements []models.Movement

	for row.Next() {
		var movement models.Movement
		err = row.Scan(&movement.Type, &transactionsDetails.ID, &movement.Date, &movement.Amount)
		if err != nil {
			log.Println(err)
			return models.Transacion_respuesta{}, err
		}
		movements = append(movements, movement)
	}
	transactionsDetails.Movements = movements

	return transactionsDetails, nil
}

func TransactionDeleteReceiverSender(person_id string) error {
	//var count int64
	sqlStatement := "DELETE FROM transaction where sender_id=$1 OR receiver_id= $1 AND type='deposit' OR type='withdrawal';"

	_, err := db.Db.Exec(sqlStatement, person_id)
	if err != nil {
		log.Println(err)
		return errors.New("No se encontro transacciones con ese documento")
	}
	/* count, err = row.RowsAffected()
	if err != nil {
		return errors.New("No se pudo eliminar las transacciones de la billetera")
	}

	if count==0{
		return errors.New("No se elimino ninguna transaccion")
	} */

	return nil

}

func TransactionUpdateTransfer(person_id string) error {
	//var count int64
	sqlStatement := "UPDATE transaction SET sender_id= NULL OR receiver_id=NULL WHERE sender_id=$1 OR receiver_id= $1;"

	_, err := db.Db.Exec(sqlStatement, person_id)
	if err != nil {
		log.Println(err)
		return errors.New("No se encontro transacciones con ese documento")
	}
	/* 	count, err = row.RowsAffected()
	   	if err != nil {
	   		return count, errors.New("No se pudo eliminar las transacciones de la billetera")
	   	}


	   	if count==0{
	   		return errors.New("No se elimino ninguna transaccion")
	   	}
	*/
	return nil

}
