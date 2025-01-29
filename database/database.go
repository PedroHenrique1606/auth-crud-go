package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	var err error
	connStr := "user=user password=password dbname=go_auth sslmode=disable host=localhost port=5432"
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Banco de dados conectado!")
}
