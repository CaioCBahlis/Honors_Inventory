package API

import (
	"Honors_Inventory/Server-Side/Equipments_DB"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

var Connection *sql.DB

type Location struct {
	Room_name     string `json:"Room_name"`
	Building_type string `json:"Building_type"`
}

func Init_DB(Connect *sql.DB) {
	Connection = Connect
}

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

func AddEquipment(w http.ResponseWriter, r *http.Request) {
	var NewEquipment Equipments_DB.Equipment

	err := json.NewDecoder(r.Body).Decode(&NewEquipment)
	if err != nil {
		log.Printf("API coudln't decode request body to Add Equipment: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = Equipments_DB.CreateEquipments(Connection, NewEquipment.Model, NewEquipment.Equipment_type, NewEquipment.Equipment_Status)
	if err != nil {
		log.Printf("API coudln't insert new Equipment: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchEquipment(w http.ResponseWriter, r *http.Request) {
	var NewEquipment Equipments_DB.Equipment

	err := json.NewDecoder(r.Body).Decode(&NewEquipment)
	if err != nil {
		log.Printf("API coudln't decode request body to Search Equipment: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res, err := Equipments_DB.SearchEquipments(Connection, NewEquipment.Model, NewEquipment.Equipment_type, NewEquipment.Equipment_Status)
	if err != nil {
		log.Printf("API coudln't retrieve searched Equipment: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	MyJson, err := json.Marshal(res)
	if err != nil {
		log.Printf("Failed to convert Searched Equipments to JSON: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(MyJson)
}

func RemoveEquipment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	err := Equipments_DB.RemoveEquipment(Connection, ID)
	if err != nil {
		log.Printf("API coudln't remove Equipment: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TransferEquipment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	var NewLocation Location

	err := json.NewDecoder(r.Body).Decode(&NewLocation)
	if err != nil {
		log.Printf("API coudln't decode request body to Location: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = Equipments_DB.EquipmentTransfer(Connection, ID, NewLocation.Room_name, NewLocation.Building_type)
	if err != nil {
		log.Printf("API coudln't transfer Equipment: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EditEquipment(w http.ResponseWriter, r *http.Request) {
	var UpdateEquipment Equipments_DB.Equipment
	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	err := json.NewDecoder(r.Body).Decode(&UpdateEquipment)
	if err != nil {
		log.Printf("API coudln't decode request body to EditEquipment: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = Equipments_DB.UpdateEquipment(Connection, ID, UpdateEquipment.Model, UpdateEquipment.Equipment_type, UpdateEquipment.Equipment_Status)
	if err != nil {
		log.Printf("API coudln't update Equipment: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<3"))
	w.WriteHeader(http.StatusOK)
}
