package Equipments_DB

import (
	"Honors_Inventory/Server-Side/Locations_DB"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Equipment struct {
	Id                    string `json:"id"`
	Model                 string `json:"model"`
	Equipment_type        string `json:"equipment_type"`
	Equipment_Status      string `json:"equipment_status"`
	Equipment_Location_ID string `json:"equipment_location_id"`
	Equipment_Room        string `json:"equipment_room"`
	Equipment_Room_Type   string `json:"equipment_room_type"`
	Inserted_at           string `json:"inserted_at"`
}

type message struct {
	Message string `json:"message"`
}

func CreateEquipments(Connection *sql.DB, model, equipment_type string) error {
	Sql_Query := "INSERT INTO equipment(model, equipment_type, equipment_status, location_id) VALUES ($1, $2, $3, $4)"

	location_id := 3              //All New Equipments default value is the storage (Foreign Key = 3);
	equipment_status := "Working" // I figured it doesn't make sense to add broken equipment to the Inventory, so default is Working

	_, err := Connection.Exec(Sql_Query, model, equipment_type, equipment_status, location_id)
	if err != nil {
		log.Printf("Fail Inserting into the Equipments table: %v", err)
		return err
	}

	UpdateAuditLogs(Connection, "Created", model+" "+equipment_type)

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
	Sql_Query := `WITH deleted AS (
 					 DELETE FROM equipment
  					WHERE id = $1
						)
			UPDATE equipment
			SET inserted_at = NOW()
			WHERE id = 1;
		`
	//This looks really weird, but is simply deletes the element at the given id, and updates the time of the first item,
	//since I'm maintaining the time of the last change, and I'm deleting an element, I need something else to keep the time
	// id=1 is my time tracker, which is not displayed in the table

	_, err := Connection.Exec(Sql_Query, id)
	if err != nil {
		log.Printf("Fail Deleting Items from Equipment Table: %v", err)
		return err
	}
	if id == 1 {
		return errors.New("Cannot Delete Time-Tracker id=1")
	}

	UpdateAuditLogs(Connection, "Removed", fmt.Sprintf("Equipment id#%d", id))

	return nil
}

func UpdateEquipment(Connection *sql.DB, id int, NewModel, NewType, NewStatus string) error {
	Sql_Query := "UPDATE equipment SET model=$1, equipment_type=$2, equipment_status=$3, inserted_at=NOW() where id=$4"
	_, err := Connection.Exec(Sql_Query, NewModel, NewType, NewStatus, id)
	if err != nil {
		log.Printf("Fail updating the Equipment: %v", err)
	}

	UpdateAuditLogs(Connection, "Updated", fmt.Sprintf("Equipment id#%d", id))

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

func EquipmentTransfer(Connection *sql.DB, id int, room_name string) error {
	Sql_Query := "UPDATE equipment SET location_id=$1, inserted_at=NOW() WHERE id=$2"

	Location_Id, err := Locations_DB.GetLocationId(Connection, room_name)
	if err != nil {
		log.Printf("Error Getting Location Id: %v", err)
		return err
	}

	_, err = Connection.Exec(Sql_Query, Location_Id, id)
	if err != nil {
		log.Printf("Fail Updating Location: %v", err)
		return err
	}

	UpdateAuditLogs(Connection, "Transfered", fmt.Sprintf("Equipment id#%d", id))

	return nil
}

func GetEquipments(Connection *sql.DB, Order string) ([]Equipment, error) {
	Sql_Query := `
    SELECT 
        e.*, 
        l.room_name, 
        l.building_type
    FROM equipment e
    JOIN locations l ON e.location_id = l.id
    WHERE e.id > 1
    ORDER BY ` + Order

	rows, err := Connection.Query(Sql_Query)
	if err != nil {
		log.Printf("Fail Selecting items from Equipment Table: %v", err)
		return nil, err
	}

	var Equipments []Equipment
	defer rows.Close()
	for rows.Next() {
		var MyEquipment Equipment
		err = rows.Scan(&MyEquipment.Id, &MyEquipment.Model, &MyEquipment.Equipment_type, &MyEquipment.Equipment_Status, &MyEquipment.Equipment_Location_ID, &MyEquipment.Inserted_at, &MyEquipment.Equipment_Room, &MyEquipment.Equipment_Room_Type)
		if err != nil {
			log.Printf("Error Scanning Rows: %v", err)
			return nil, err
		}

		Equipments = append(Equipments, MyEquipment)
	}

	return Equipments, nil
}

func GetEquipmentsInfo(Connection *sql.DB) ([]int, error) {
	MyStatsQuery := `
	SELECT
		COUNT (id) AS total,
		COUNT(CASE WHEN location_id=3 THEN 1 END) AS Warehouse,
		COUNT(CASE WHEN equipment_status IN ('Needs Maintenance', 'Broken') THEN 1 END) as Maintenance
	FROM equipment;
    `

	var total, Warehouse, Maintenance int
	rows := Connection.QueryRow(MyStatsQuery)

	err := rows.Scan(&total, &Warehouse, &Maintenance)
	if err != nil {
		log.Printf("Error Scanning Equipment Info: %v", err)
		return nil, err
	}

	return []int{total, Warehouse, Maintenance}, nil
}

func GetLastInsertion(Connection *sql.DB) (string, error) {
	My_Query := `SELECT inserted_at,
       to_char(
    inserted_at AT TIME ZONE 'America/New_York',
    'YYYY-MM-DD HH24:MI:SS TZ')
	 AS insert_at_tampa FROM equipment
		ORDER BY inserted_at DESC ;
	`

	var val string
	var time string

	err := Connection.QueryRow(My_Query).Scan(&val, &time)
	if err != nil {
		log.Printf("Error Scanning Time rows: %v", err)
		return "", err
	}

	return time, err
}

func GetAuditLogs(Connection *sql.DB) ([]message, error) {
	MyQuery := `SELECT message FROM audit ORDER BY id DESC limit 10 `

	rows, err := Connection.Query(MyQuery)
	if err != nil {
		log.Printf("Error Getting audit logs: %v", err)
		return nil, err
	}
	defer rows.Close()

	var MyLogs []message
	for rows.Next() {
		var MyMessage message
		err = rows.Scan(&MyMessage.Message)
		if err != nil {
			log.Printf("Error Scanning Audit Logs Rows: %v", err)
			return nil, err
		}
		MyLogs = append(MyLogs, MyMessage)
	}

	return MyLogs, nil
}

func UpdateAuditLogs(Connection *sql.DB, OP_Type string, Data string) {
	Mylocation, _ := time.LoadLocation("America/New_York")
	TampaTime := time.Now().In(Mylocation)
	PrettyTime := TampaTime.Format("2006-01-02 15:04:05")

	msg := fmt.Sprintf("Rocky The Bull %s  %s at %v", OP_Type, Data, PrettyTime)

	MyQuery := `INSERT INTO audit(message) VALUES($1)`
	_, err := Connection.Exec(MyQuery, msg)
	if err != nil {
		log.Printf("Fail Updating AuditLogs: %v", err)
	}
}
