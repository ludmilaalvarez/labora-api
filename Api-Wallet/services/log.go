package services

import (
	"Api-Wallet/db"
	"Api-Wallet/models"
	"bytes"
	"encoding/json"

	//	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type PostgresLog struct {
	Db db.DbConnection
}

var Idwallet int

func Request(Person_id string, country string) (models.Respuesta, error) {
	var client = &http.Client{}
	var nuevaRespuesta models.Respuesta

	API := "https://api.checks.truora.com/v1/checks/"
	TOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k4YWJkOWE1ZGFmNzM1NGQ1YjVlZjVjYTI4MjJhMjA3OSIsImV4cCI6MzI2MTY4OTIwMiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ4ODkyMDIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX3hUSGxqU1d2RCIsImp0aSI6IjM2YTZiNGJlLTM3NTUtNGQzMC04ZTM0LTNmZDMyOGI3ZDk3NCIsImtleV9uYW1lIjoidHJ1Y29kZSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoidHJ1b3JhdGVhbW5ld3Byb2QtdHJ1Y29kZSJ9.PuE6cS6938PbQz_4qMLySs9dr3fywFqqGdfcF6Suw0U"

	body, _ := json.Marshal(map[string]string{
		"national_id":     Person_id,
		"country":         country,
		"type":            "person",
		"user_authorized": "true",
	})

	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, API, payload)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Truora-API-Key", TOKEN)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &nuevaRespuesta)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return nuevaRespuesta, nil
}

func (p *PostgresLog) CrearSolicitud(Datos *models.Datos_Solicitados) (string, error) {
	var (
		status    string
		id        int
		solicitud models.Solicitud
	)

	status, err := VerificarStatusScore(Datos)

	solicitud = models.Solicitud{
		Id:               id,
		Person_id:        Datos.National_id,
		Date:             time.Now(),
		Country:          Datos.Country,
		Wallet_id:        Idwallet,
		Status:           status,
		State:            Datos.State,
		Type_transaction: "Create Wallet",
	}

	insertStatement := `INSERT INTO solicitud (state, date, status, person_id, country, wallet_id, type_transaction)
                        VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Db.Exec(insertStatement, solicitud.State, solicitud.Date, solicitud.Status, solicitud.Person_id, solicitud.Country, solicitud.Wallet_id, solicitud.Type_transaction)
	if err != nil {
		log.Println(err)
	}
	return status, nil

}

func VerificarStatusScore(Datos *models.Datos_Solicitados) (string, error) {
	var (
		status    string
		comprobar bool
	)

	Datos_obtenidos, err := Request(Datos.National_id, Datos.Country)

	if Datos_obtenidos.Check.Summary.NamesFound != nil {

		nombre := Datos_obtenidos.Check.Summary.NamesFound[0]
		str := fmt.Sprintln(nombre.FirstName, nombre.LastName)
		comprobar = (strings.ToUpper(Datos.Name) == strings.TrimSpace(strings.ToUpper(str)))

	} else {

		return ("\nDatos no validos!"), err
	}

	if (Datos_obtenidos.Check.Score == 1) && comprobar {
		existencia := ComprobarExistenciaWallet(Datos.National_id)

		if existencia {
			status = "Completado"
			Idwallet, err = WalletHandler.CrearWallet(Datos)
		} else {
			return ("\nYa existe una billetera con ese Documento."), err
		}
	} else {
		status = "Denegado"
	}

	return status, nil

}

func RegistrarTransaccion(status string, newTransaccion models.Transaction) error {

	wallet_id, country, state := BuscarIDWallet(newTransaccion.Sender_id)

	insertStatement := `INSERT INTO solicitud (state, date, status, person_id, country, wallet_id, type_transaction)
                        VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Db.Exec(insertStatement, state, time.Now(), status, newTransaccion.Sender_id, country, wallet_id, newTransaccion.Type)
	if err != nil {
		log.Println(err)
		return err
	}

	if newTransaccion.Type == "Transfer" {
		wallet_id, country, state := BuscarIDWallet(newTransaccion.Receiver_id)

		insertStatement := `INSERT INTO solicitud (state, date, status, person_id, country, wallet_id, type_transaction)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

		_, err := db.Db.Exec(insertStatement, state, time.Now(), status, newTransaccion.Receiver_id, country, wallet_id, "Transfer received")
		if err != nil {
			log.Println(err)
			return err
		}

	}

	return nil

}
