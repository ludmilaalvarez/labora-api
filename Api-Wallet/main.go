package main

import (
	"Api-Wallet/controllers"
	"Api-Wallet/db"
	"Api-Wallet/services"

	//"Api-Wallet/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	db.EstablishDbConnection()
	services.Init()

	router := mux.NewRouter()

	router.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/WalletStatus/{national_id}", controllers.StatusWallet).Methods("GET")
	router.HandleFunc("/Transaction", controllers.Transaction).Methods("POST")
	router.HandleFunc("/TransactionHistory/{wallet_id}", controllers.TransactionHistory).Methods("GET")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8000", "http://example.com"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Agregar el middleware CORS a todas las rutas
	handler := corsOptions.Handler(router)

	http.ListenAndServe(":3000", handler)
}
