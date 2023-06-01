package controllers

import (
	//	"Api-Wallet/db"
	"Api-Wallet/models"
	"Api-Wallet/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	//	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	var Datos models.Datos_Solicitados
	var resultado string

	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	json.Unmarshal(rqBody, &Datos)

	resultado, err = services.LogHandler.CrearSolicitud(&Datos)

	if resultado != "Completado" {
		w.Write([]byte("Error al crear la billetera"))
		w.Write([]byte(resultado))

	} else {
		w.Write([]byte("La billetera ha sido creada con exito!"))
	}

}

func StatusWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	Dni := params["national_id"]

	wallet, err := services.WalletHandler.StatusWallet(string(Dni))

	if err != nil {
		fmt.Fprintf(w, "Inserte un item valido")
		return
	}
	json.NewEncoder(w).Encode(wallet)

}

func Transaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTransaccion models.Transaction

	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos!")
		return
	}
	json.Unmarshal(rqBody, &newTransaccion)
	fmt.Println(newTransaccion)

	err = services.CreateTransaction(newTransaccion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransaccion)

}

func TransactionHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	wallet_id_raw := params["wallet_id"]
	wallet_id, err := strconv.Atoi(wallet_id_raw)
	transacciones, err := services.HistorialTransacciones(wallet_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transacciones)
}
