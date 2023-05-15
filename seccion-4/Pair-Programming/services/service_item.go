package services

import (
	"Pair-Programming/models"
	"database/sql"

	"fmt"
	"log"
)


func Get(tipo string, dato string) ([]models.Item, error) {

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
		rows, err = Db.Query(query, dato)
	} else {
		rows, err = Db.Query(query)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.ItemDetails)
		if err != nil {
			fmt.Println(err)
			continue
		}
		item.PrecioTotal()
		
		items = append(items, item)
	}
	return items, nil
}

func GetItemsPorPagina(pageIndex int, itemsPerPageInt int)([]models.Item, error){
	var items = make([]models.Item, 0)

	init := itemsPerPageInt * (pageIndex - 1)
	limit := itemsPerPageInt

	query:="SELECT * FROM items LIMIT $1 OFFSET $2"
	rows, err := Db.Query(query, limit, init)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.ItemDetails)
		if err != nil {
			log.Println(err)
		}
		item.PrecioTotal()
		
		items = append(items, item)
	}
	return items, nil
}


func CreateNewItem(item models.Item)error {

	insertStatement := `INSERT INTO items (customer_name, order_date, product, quantity, price, details)
                        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := Db.Exec(insertStatement, item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, item.ItemDetails)
	return err

}

func UpdateItem(id int, item models.Item)( int64, error){

	updateStatement:="UPDATE items SET customer_name=$1, order_date=$2,product=$3, quantity=$4, price=$5, details=$6 WHERE id=$7 "
	row, err:= Db.Exec(updateStatement, item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, item.ItemDetails, id )

	if err != nil{
		fmt.Println(err)
	}

    count, err := row.RowsAffected()
	  if err != nil {
		panic(err)
	}
	return count, err

}

func DeleteItem(id int)int64{

	sqlStatement := "DELETE FROM items WHERE id = $1;"

	row, err := Db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	}
	count, err := row.RowsAffected()
	if err != nil {
	  panic(err)
    }
	return count

}


