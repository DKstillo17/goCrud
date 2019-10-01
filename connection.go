package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

//Conexion a BD

func getConnection() *sql.DB {
	dsn := "postgres://postgres:admin@127.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
