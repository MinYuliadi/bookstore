package config

import (
	"bookstore/migrations"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("cannot open database", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("cannot connect to database", err.Error())
	}

	DB = db
	fmt.Println("connected to postgres database")

	migrations.DBMigrate(db)
}
