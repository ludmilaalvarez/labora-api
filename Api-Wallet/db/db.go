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

var (
	host     = os.Getenv("HostDB")
	port     = os.Getenv("PortDB")
	dbName   = os.Getenv("DbName")
	user     = os.Getenv("UserDB")
	password = os.Getenv("PasswordDB")
)

var dbConn *sql.DB

func EstablishDbConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	psqlInfo :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conexion", dbConn)
	Db = DbConnection{dbConn}
}
