package main

import (
	"Honors_Inventory/Server-Side/API"
	"database/sql"
	"github.com/go-chi/chi/v5"
)

func RouterSetup(r *chi.Mux, db *sql.DB) {
	API.Init_DB(db)

	r.Get("/Heartbeat", API.Heartbeat)
	r.Get("/API/GetMaintenance", API.GetMaintenanceEquipment)

	r.Post("/API/AddEquipment", API.AddEquipment)
	r.Post("/API/SearchEquipment", API.SearchEquipment)

	r.Put("/API/TransferEquipment/{id}", API.TransferEquipment)
	r.Put("/API/EditEquipment/{id}", API.EditEquipment)

	r.Delete("/API/RemoveEquipment/{id}", API.RemoveEquipment)
}
