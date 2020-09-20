package services

import (
	"encoding/json"
	"net/http"

	"go-crud-api/contracts"

	"github.com/gorilla/mux"
)

// GetAnEvent - Responsible for listing a specific event
func GetAnEvent(w http.ResponseWriter, r *http.Request) {
	var response contracts.Event

	eventID := mux.Vars(r)["id"]
	allEvents := contracts.Events

	for _, event := range allEvents {
		if event.ID != eventID {
			continue
		}

		response = event
	}

	json.NewEncoder(w).Encode(response)
}
