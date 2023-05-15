package controllers

import (
	"Pair-Programming/services"
	"Pair-Programming/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := services.Get("items", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(items)
}
func GetItemsPorPagina(w http.ResponseWriter, r *http.Request) {
	
	params:=r.URL.Query()
	page:= params["page"]
	itemsPerPage:=params["itemsPerPage"]

	pageIndex, _ := strconv.Atoi(page[0])
	itemsPerPageInt, _ := strconv.Atoi(itemsPerPage[0])

	items, err := services.GetItemsPorPagina(pageIndex, itemsPerPageInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(items)
}
func GetItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	
	item, err := services.Get("id", string(id))

	if err != nil {
		fmt.Fprintf(w, "Inserte un item valido")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func GetName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	name:= params["name"]
	item, err:= services.Get("name", name)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item valido")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item

	rqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}
	
	json.Unmarshal(rqBody, &newItem)
	
	err= services.CreateNewItem(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("¡Item creado exitosamente!")
}


func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var item models.Item

	json.NewDecoder(r.Body).Decode(&item)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El ID debe ser un número"))
		return
	}
	
	count , err:= services.UpdateItem(id, item)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al actualizar el item"))
		
	}
	//Se fina si hubo algun cambio la base de datos
	if count==0{
		err= services.CreateNewItem(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("¡Item creado exitosamente!")
	}
	json.NewEncoder(w).Encode(item)
	
} 

func DeleteItem(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	var item models.Item

	json.NewDecoder(r.Body).Decode(&item)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El ID debe ser un número"))
		return
	}

	count:= services.DeleteItem(id)
	if count>0{
		fmt.Println("Item Eliminado")
		w.Write([]byte("Item Eliminado con suceso!"))
	
	}
}
