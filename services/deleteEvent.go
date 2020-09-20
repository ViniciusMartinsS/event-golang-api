package services

import (
	"encoding/json"
	"go-crud-api/contracts"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteEvent - Responsible by deleting a specific event from database
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, event := range contracts.Events {
		if event.ID != eventID {
			continue
		}

		deletedEvent := event
		contracts.Events = append(contracts.Events[:i], contracts.Events[i+1:]...)
		json.NewEncoder(w).Encode(deletedEvent)
	}
}
