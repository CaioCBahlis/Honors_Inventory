package Equipments_DB

import (
	"Honors_Inventory/Server-Side/Locations_DB"
	"database/sql"
	"log"
)

type Equipment struct {
	Id                    int
	Model                 string
	Equipment_type        string
	Equipment_Status      string
	Equipment_Location_ID int
}

func CreateEquipments(Connection *sql.DB, model, equipment_type, equipment_status string) error {
	Sql_Query := "INSERT INTO equipment(model, equipment_type, equipment_status, location_id) VALUES ($1, $2, $3, $4)"

	location_id := 1 //All New Equipments default value is the storage (Foreign Key = 1);

	_, err := Connection.Exec(Sql_Query, model, equipment_type, equipment_status, location_id)
	if err != nil {
		log.Printf("Fail Inserting into the Equipments table: %v", err)
		return err
	}
	return nil
}

func SearchEquipments(Connection *sql.DB, model, equipment_type, equipment_status string) ([]Equipment, error) {
	Sql_Query := "SELECT * FROM equipment WHERE model = $1 AND equipment_type = $2 AND equipment_status = $3"

	rows, err := Connection.Query(Sql_Query, model, equipment_type, equipment_status)
	if err != nil {
		log.Printf("Fail Selecting items from the Locations Table: %v", err)
	}
	var Equipments []Equipment
	defer rows.Close()
	for rows.Next() {
		var MyEquipment Equipment
		err = rows.Scan(&MyEquipment.Id, &MyEquipment.Model, &MyEquipment.Equipment_type, &MyEquipment.Equipment_Status, &MyEquipment.Equipment_Location_ID)
		if err != nil {
			log.Printf("Error Scanning Rows: %v", err)
		}
		Equipments = append(Equipments, MyEquipment)
	}
	return Equipments, nil
}

func RemoveEquipment(Connection *sql.DB, id int) error {
	Sql_Query := "DELETE FROM equipment WHERE id=$1"

	_, err := Connection.Exec(Sql_Query, id)
	if err != nil {
		log.Printf("Fail Deleting Items from Equipment Table: %v", err)
		return err
	}
	return nil
}

func UpdateEquipment(Connection *sql.DB, id int, NewModel, NewType, NewStatus string) error {
	Sql_Query := "UPDATE equipment SET model=$1, equipment_type=$2, equipment_status=$3 where id=$4"
	_, err := Connection.Exec(Sql_Query, NewModel, NewType, NewStatus, id)
	if err != nil {
		log.Printf("Fail updating the Equipment: %v", err)
	}

	return nil
}

func EquipmentsForMaintenace(Connection *sql.DB) ([]Equipment, error) {
	Sql_Query := "Select * FROM equipment WHERE equipment_status NOT LIKE 'working'"
	rows, err := Connection.Query(Sql_Query)
	if err != nil {
		log.Printf("Fail Getting Items that need Maitenance: %v", err)
	}
	defer rows.Close()

	var ForMaintenace []Equipment
	for rows.Next() {
		var MyEquipment Equipment
		err = rows.Scan(&MyEquipment.Id, &MyEquipment.Model, &MyEquipment.Equipment_type, &MyEquipment.Equipment_Status, &MyEquipment.Equipment_Location_ID)
		if err != nil {
			log.Printf("Error Scanning Rows: %v", err)
		}
		ForMaintenace = append(ForMaintenace, MyEquipment)
	}
	return ForMaintenace, nil
}

func EquipmentTransfer(Connection *sql.DB, id int, room_name, building_type string) error {
	Sql_Query := "UPDATE equipment SET location_id=$1 WHERE id=$2"

	Location_Id, err := Locations_DB.GetLocationId(Connection, room_name, building_type)
	if err != nil {
		log.Printf("Error Getting Location Id: %v", err)
		return err
	}

	_, err = Connection.Exec(Sql_Query, Location_Id, id)
	if err != nil {
		log.Printf("Fail Updating Location: %v", err)
		return err
	}
	return nil
}
