package controllers

import (
	"Api-Wallet/db"
	"Api-Wallet/models"
	"Api-Wallet/services"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var logService services.LogService
var walletService services.WalletService

func Init() {
	walletService = services.WalletService{
		DbHandlers: &services.PostgresWallet{Db: db.Db},
	}
	logService = services.LogService{
		DbHandlers: &services.PostgresLog{Db: db.Db},
	}
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	Init()
	var Datos models.Datos_Solicitados
	var resultado string

	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	json.Unmarshal(rqBody, &Datos)

	resultado = logService.CrearSolicitud(&Datos)

	if resultado != "Completado" {
		w.Write([]byte("Error al crear la billetera"))
		w.Write([]byte(resultado))

	} else {
		w.Write([]byte("La billetera ha sido creada con exito!"))
	}

}

func StatusWallet(w http.ResponseWriter, r *http.Request) {
	Init()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	Dni := params["national_id"]

	wallet, err := walletService.StatusWallet(string(Dni))

	if err != nil {
		fmt.Fprintf(w, "Inserte un item valido")
		return
	}
	json.NewEncoder(w).Encode(wallet)

}
