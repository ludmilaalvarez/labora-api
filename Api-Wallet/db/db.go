package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConnection struct {
	*sql.DB
}

var Db DbConnection

const (
	host   = "localhost"
	port   = 5432
	dbName = "labora-wallet"
	user   = "postgres"
)

var dbConn *sql.DB

func EstablishDbConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	password := os.Getenv("PasswordDB")
	psqlInfo :=
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conexion", dbConn)
	Db = DbConnection{dbConn}
}
