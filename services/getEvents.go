package services

import (
	"encoding/json"
	"go-crud-api/contracts"
	"net/http"
)

// GetEvents - Responsible for getting all saved events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contracts.Events)
}
