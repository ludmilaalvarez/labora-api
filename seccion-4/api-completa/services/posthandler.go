package services

import (
	"Pair-Programming/models"
	"database/sql"

	"fmt"
	//"log"
	//"sync"
)


type PostgresDBHandler struct {
    Db *sql.DB
}

//var Mutex sync.Mutex

func (p *PostgresDBHandler)Get(tipo string, dato string)([]models.Item, error){
	var items = make([]models.Item, 0)
	var query string

	switch tipo {
	case "items":
		query = "SELECT * FROM items"
	case "id":
		query = "SELECT * FROM items WHERE id=$1"
	case "name":
		query = "SELECT * FROM items WHERE customer_name=$1"
	}

	var rows *sql.Rows
	var err error

	if dato != "" {
		rows, err = p.Db.Query(query, dato)
	} else {
		rows, err = p.Db.Query(query)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		
		err := rows.Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.ItemDetails, &item.View)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//item.VistasTotal()
		//VistasContador(item.View, item.Id)
		item.PrecioTotal()
		
		items = append(items, item)
	}
	return items, nil
}

/* func (p *PostgresDBHandler) CreateItem(item Item) error {
    // Implementar la lógica para crear un artículo en la base de datos PostgreSQL
}

func (p *PostgresDBHandler) GetItem(id int) (Item, error) {
    // Implementar la lógica para obtener un artículo de la base de datos PostgreSQL
}

func (p *PostgresDBHandler) UpdateItem(item Item) error {
    // Implementar la lógica para actualizar un artículo en la base de datos PostgreSQL
}

func (p *PostgresDBHandler) DeleteItem(id int) error {
    // Implementar la lógica para eliminar un artículo de la base de datos PostgreSQL
}

 */