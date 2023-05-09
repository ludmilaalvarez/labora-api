package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Item struct {
	Id           int    //`json:"id"`
	CustomerName string //`json:"name"`
	OrderDate    string //`json:"string"`
	Product      string //`json:"product"`
	Quantity     int    // `json:"quantity"`
	Price        float64
	ItemDetails  string //`json:"price"`
}

const (
	host     = "localhost"
	port     = 5432
	dbName   = "labora-proyect-1"
	user     = "postgres"
	password = "admin"
)

var ErrInvalidInput = errors.New("entrada inválida")
var items []Item

func main() {

	establishDbConnection()

	router := mux.NewRouter()

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/porID/{id}", getItem).Methods("GET")
	router.HandleFunc("/items/porNombre/{name}", getName).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	//router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":3000", router)

}

func establishDbConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conexion", dbConn)
	//defer dbConn.Close()
	return dbConn, err
}

func getItems(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	dbConn, err := establishDbConnection()

	//fmt.Println("conexion", dbConn)
	query := `select * from items`

	rows, err := dbConn.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.ItemDetails)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	dbConn, err := establishDbConnection()

	params := mux.Vars(r)

	//Convierte una cadena en entero
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Inserte un item valido")
		return
	}

	var item Item
	err = dbConn.QueryRow("SELECT * FROM items WHERE id=$1", id).Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.ItemDetails)
	if err != nil {
		fmt.Fprintf(w, "Error getting item from database")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func getName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbConn, err := establishDbConnection()

	// Función para obtener un elemento específico
	parametros := mux.Vars(r)
	name := parametros["name"]

	var item Item
	err = dbConn.QueryRow("SELECT * FROM items where customer_name=$1", name).Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.ItemDetails)
	if err != nil {
		fmt.Fprintf(w, "Error getting item from database")
		return
	}
	json.NewEncoder(w).Encode(item)

}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	json.Unmarshal(rqBody, &newItem)

	fmt.Println(newItem)

	dbConn, err := establishDbConnection()

	// Insertar el nuevo item en la base de datos
	insertStatement := `INSERT INTO items (customer_name, order_date, product, quantity, price, details)
                        VALUES ($1, Date($2), $3, $4, $5, $6)`
	_, err = dbConn.Exec(insertStatement, newItem.CustomerName, newItem.OrderDate, newItem.Product, newItem.Quantity, newItem.Price, newItem.ItemDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar una respuesta exitosa al cliente
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("¡Item creado exitosamente!")
}

/* func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	dbConn, err := establishDbConnection.Prepare(updateStatement)

	var nuevonombre string
	var encontrado bool

	json.NewDecoder(r.Body).Decode(&nuevonombre)

	updateStatement := `UPDATE items SET customer_name=? WHERE id=?)`

	for index, item := range items {
		if item.Id == params["id"] {
			dbConn.Prepare(updateStatement)

_, err = dbConn.Exec(insertStatement, newItem.CustomerName, newItem.OrderDate, newItem.Product, newItem.Quantity, newItem.Price, newItem.ItemDetails)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
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
} */

func deleteItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	dbConn, _ := establishDbConnection()

	// Función para obtener un elemento específico
	parametros := mux.Vars(r)
	id := parametros["id"]

	var item Item
	sqlStatement := `
		DELETE FROM items
		WHERE id = $1;`

	_, err := dbConn.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(item)
}
