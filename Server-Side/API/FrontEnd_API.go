package API

import (
	"Honors_Inventory/Server-Side/Equipments_DB"
	"encoding/json"
	"log"
	"net/http"
)

func GetMaintenanceEquipment(w http.ResponseWriter, r *http.Request) {
	Items, err := Equipments_DB.EquipmentsForMaintenace(Connection)
	if err != nil {
		log.Printf("API coudln't retrieve items for maitenance: %v\n", err)
		http.Error(w, http.StatusText(500), 500)
	}

	MyJson, err := json.Marshal(Items)
	if err != nil {
		log.Printf("Failed to convert Maintenace Equipments to JSON: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(MyJson)
}

func GetEquipments(w http.ResponseWriter, r *http.Request) {
	Equipments, err := Equipments_DB.EquipmentsForMaintenace(Connection)
	if err != nil {
		log.Printf("API coudln't retrieve items for maintenace: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	MyJson, err := json.Marshal(Equipments)
	if err != nil {
		log.Printf("Failed to convert Equipments to JSON: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(MyJson)
}

func GetEquipmentInfo(w http.ResponseWriter, r *http.Request) {
	Equipment_Info, err := Equipments_DB.GetEquipmentsInfo(Connection)
	if err != nil {
		log.Printf("Failed to retrieve Equipment Info: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	MyJson, err := json.Marshal(Equipment_Info)
	if err != nil {
		log.Printf("Failed to Marshall equipment info to JSON: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(MyJson)
}
