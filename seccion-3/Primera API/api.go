package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

var ErrInvalidInput = errors.New("entrada inválida")

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

//	router.HandleFunc("/", getUno)
	router.HandleFunc("/items", get("items")).Methods("GET")
	router.HandleFunc("/items/porID/{id}", get("id")).Methods("GET")
	router.HandleFunc("/items/details", getDetails).Methods("GET")
	router.HandleFunc("/items/porNombre/{name}", getName).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}

/* func getUno(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("Mi primera Api"))
}

func getItems(w http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(items)
	if err != nil {
		fmt.Println(ErrInvalidInput)
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

func getDetails(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	wg := &sync.WaitGroup{}
	detailsChannel := make(chan ItemDetails, len(items))
	var detailedItems []ItemDetails
	for _, item := range items {
		wg.Add(1)

		go func(id string) {
			itemDetails := getItemDetails(id)
			defer wg.Done()
			detailsChannel <- itemDetails

		}(item.ID)
	}
	wg.Wait()
	close(detailsChannel)
	for details := range detailsChannel {
		detailedItems = append(detailedItems, details)
	}
	json.NewEncoder(w).Encode(detailedItems)

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

} */

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

func getItemDetails(id string) ItemDetails {

	if _, err := strconv.Atoi(id); err != nil {
		fmt.Println(ErrInvalidInput)
	}

	time.Sleep(100 * time.Millisecond)
	var foundItem Item
	for _, item := range items {
		if item.ID == id {
			foundItem = item
			break
		}
	}
	return ItemDetails{
		Item:    foundItem,
		Details: fmt.Sprintf("Detalles para el item %s", id),
	}
}

func getGeneral(w http.ResponseWriter, r *http.Request)

func get( tipo string, w http.ResponseWriter, r *http.Request){
	switch tipo{
	case "items":
		json.NewEncoder(w).Encode(items)
	case "item":
		params := mux.Vars(r)
		id:= params["id"]
		json.NewEncoder(w).Encode(buscar(id))

	case "details":
		wg := &sync.WaitGroup{}
		detailsChannel := make(chan ItemDetails, len(items))
		var detailedItems []ItemDetails
		for _, item := range items {
			wg.Add(1)
	
			go func(id string) {
				itemDetails := getItemDetails(id)
				defer wg.Done()
				detailsChannel <- itemDetails
	
			}(item.ID)
		}
		wg.Wait()
		close(detailsChannel)
		for details := range detailsChannel {
			detailedItems = append(detailedItems, details)
		}
		json.NewEncoder(w).Encode(detailedItems)

	case "name":
		params := mux.Vars(r)
		name:= params["name"]
		json.NewEncoder(w).Encode(buscar(name))
	
	}

}

func buscar( dato string) []Item {
	var encontrado bool = false
	for _, item := range items {
		if ( strings.ToLower(item.Name) == strings.ToLower(dato) ) || (item.ID == dato) {
			var returnable = []Item{item}
			encontrado = true
			return returnable
		}
	}
	if !encontrado {
		return items
	}

	return items
}
