package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	// Conexión a la base de datos
	cnn, err := sql.Open("mysql", "sql5741567:JQvbkpQU39@tcp(sql5.freesqldatabase.com:3306)/sql5741567")
	if err != nil {
		log.Println("[ERROR]: Unable open connection to DB: ", err.Error())
	}
	log.Println("[INFO]: Successfully connected to DB")

	db = cnn
	Ping()
	return db
}

// Exec this part of code in case your DB is empty
func MakeMigrations() error {
	db := GetConnection()
	query := `CREATE TABLE IF NOT EXISTS users (
	        email VARCHAR(64) PRIMARY KEY,
       		password VARCHAR(100) NOT NULL,
       		username VARCHAR(64) NOT NULL,
	        phone VARCHAR(10) NOT NULL
	      );`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func Ping() {
	log.Println("Test db connection")
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
