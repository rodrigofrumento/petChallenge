package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnDb() {
	// Lê .env
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL não encontrada")
	}

	// Abre conexão com banco
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Teste de conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Define variável DB
	DB = db

}
