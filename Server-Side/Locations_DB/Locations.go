package Locations_DB

import (
	"database/sql"
	"log"
)

func CreateLocations(Connection *sql.DB, room_name, building_type string) error {
	Sql_Query := "INSERT INTO locations(room_name, building_type) VALUES ($1, $2);"

	_, err := Connection.Exec(Sql_Query, room_name, building_type)

	if err != nil {
		log.Printf("Fail Creating a New Location table: %v", err)
		return err
	}
	return nil
}

func GetLocationId(Connection *sql.DB, room_name, building_type string) (int, error) {
	Sql_Query := "SELECT id FROM locations WHERE room_name = $1 AND building_type = $2;"

	var id int
	rows, err := Connection.Query(Sql_Query, room_name, building_type)
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
