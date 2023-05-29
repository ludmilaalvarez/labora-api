package main

import (
	"Api-Wallet/controllers"
	"Api-Wallet/db"

	//"Api-Wallet/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db.EstablishDbConnection()

	router := mux.NewRouter()

	router.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/WalletStatus/{national_id}", controllers.StatusWallet).Methods("GET")

	http.ListenAndServe(":3000", router)
}
