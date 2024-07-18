package controllers

import (
	"encoding/json"
	"go-template/config"
	"go-template/models"
	"net/http"
	// Add any other imports required by the GetItems function
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	items := []models.Item{}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
