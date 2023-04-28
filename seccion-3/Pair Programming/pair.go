package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

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
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}

func getUno(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("Mi primera Api"))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query() // page - itemsPerPage
	page := params["page"]
	itemsPerPage := params["itemsPerPage"]
	if len(page) == 0 {
		page = append(page, "1")
	}
	if len(itemsPerPage) == 0 {
		itemsPerPage = append(itemsPerPage, "3")
	}
	pageIndex, _ := strconv.Atoi(page[0])
	itemsPerPageInt, _ := strconv.Atoi(itemsPerPage[0])
	var newListItems []Item
	init := itemsPerPageInt * (pageIndex - 1)
	limit := init + itemsPerPageInt
	nroPage := float64(len(items)) / float64(itemsPerPageInt)
	nroPage = math.Ceil(nroPage)
	if pageIndex <= int(nroPage) {
		if limit > len(items) {
			newListItems = items[init:]
		} else {
			newListItems = items[init:limit]
		}
	}
	json.NewEncoder(w).Encode(newListItems)
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

func createItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = params["id"]
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
