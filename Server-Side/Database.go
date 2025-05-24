package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func ConnectToDatabase() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for i := 0; i < 10; i++ {
		Connection, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Println("Failed to connect to database: %v", err)
			log.Println("Retrying in 1 second...")

		}

		err = Connection.Ping()
		if err != nil {
			log.Println("Failed to connect to database: %v", err)
			log.Println("Retrying in Now...")
		}
		time.Sleep(1 * time.Second)
	}

	log.Println("Successfully connected to database :)")
}
