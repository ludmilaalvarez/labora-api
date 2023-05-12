package main

import (
	"Pair-Programming/controllers"
	"Pair-Programming/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	services.EstablishDbConnection()

	router := mux.NewRouter()

	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	router.HandleFunc("/items/porID/{id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/items/porNombre/{name}", controllers.GetName).Methods("GET")
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
	http.ListenAndServe(":3000", router)

}
