package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectToDatabase() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Connection_String := fmt.Sprintf(
		"host=localhost "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DB"),
	)

	Connection, err = sql.Open("postgres", Connection_String)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = Connection.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database :)")
}

func InsertIntoDatabase(room_name, building_type string) {
	Sql_Query := "INSERT INTO locations(room_name, building_type) VALUES ($1, $2)"
	res, err := Connection.Exec(Sql_Query, room_name, building_type)
	if err != nil {
		log.Fatalf("Failed to insert into locations: %v", err)
	}
	fmt.Println(res)
}
