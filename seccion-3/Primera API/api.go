package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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
	router.HandleFunc("/items/porNombre/{name}", getName).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}

func getUno(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("Mi primera Api"))
}

func getItems(w http.ResponseWriter, r *http.Request) {

	/* b, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}
	w.Write(b) */
	json.NewEncoder(w).Encode(items)
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

func getName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var encontrado bool = false
	for _, item := range items {
		if strings.ToLower(item.Name) == strings.ToLower(params["name"]) {
			json.NewEncoder(w).Encode(item)
			encontrado = true
		}
	}
	if !encontrado {
		json.NewEncoder(w).Encode(&Item{})
	}

}

func createItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Item
	json.Unmarshal(reqBody, &item)
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var nuevonombre string
	var encontrado bool
	json.NewDecoder(r.Body).Decode(&nuevonombre)
	for index, item := range items {
		if item.ID == params["id"] {
			items[index].Name = nuevonombre
			encontrado = true
			break
		}
	}
	if !encontrado {
		var nuevoitem Item
		nuevoitem = Item{ID: params["id"], Name: nuevonombre}
		items = append(items, nuevoitem)
	}
	json.NewEncoder(w).Encode(items)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
