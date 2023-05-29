package services

import (
	"Pair-Programming/models"
)


type DBHandler interface {
	Get(tipo string, dato string) ([]models.Item, error)
    //CreateItem(item Item) error
    //GetItem(id int) (Item, error)
    //UpdateItem(item Item) error
   // DeleteItem(id int) error
}