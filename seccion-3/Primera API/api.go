package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Mauro"},
	{ID: "2", Name: "Pedro"},
	{ID: "3", Name: "Juan"},
	{ID: "4", Name: "Candela"},
	{ID: "5", Name: "Agostina"},
	{ID: "6", Name: "Teo"},
	{ID: "7", Name: "Gala"},
	{ID: "8", Name: "Roberto"},
	{ID: "9", Name: "Cristian"},
	{ID: "10", Name: "Abril"}}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", getUno)
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	http.ListenAndServe(":3000", router)
}

func getUno(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("Mi primera Api"))
}

func getItems(w http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}
	w.Write(b)
	//json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}
