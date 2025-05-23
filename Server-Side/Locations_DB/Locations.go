package Locations_DB

import (
	"database/sql"
	"log"
)

type Locations struct {
	Room_Name     string
	building_type string
}

func CreateLocations(Connection *sql.DB, room_name, building_type string) error {
	Sql_Query := "INSERT INTO locations(room_name, building_type) VALUES ($1, $2);"

	_, err := Connection.Exec(Sql_Query, room_name, building_type)

	if err != nil {
		log.Printf("Fail Creating a New Location table: %v", err)
		return err
	}
	return nil
}

func GetLocationId(Connection *sql.DB, room_name string) (int, error) {
	Sql_Query := "SELECT id FROM locations WHERE room_name = $1;"

	var id int
	rows, err := Connection.Query(Sql_Query, room_name)
	if err != nil {
		log.Printf("Fail retrieving Location_ID from database: %v", err)
		return -1, err
	}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("Fail Scanning Location_ID: %v", err)
			return -1, err
		}
	}

	return id, nil

}

func GetLocationFromID(Connection *sql.DB, id string) (Locations, error) {
	Sql_Query := "SELECT (room_name, building_type) FROM locations WHERE id = $1;"
	var MyLocation Locations

	err := Connection.QueryRow(Sql_Query, id).Scan(&MyLocation.Room_Name, &MyLocation.building_type)
	if err != nil {
		log.Printf("Fail Retrieving Location from database: %v", err)
		return Locations{}, err
	}

	return MyLocation, nil
}
