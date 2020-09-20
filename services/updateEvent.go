package services

import (
	"encoding/json"
	"go-crud-api/contracts"
	"go-crud-api/utils"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// UpdateEvent - Responsible by updating a specific event
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var updateEvent contracts.Event

	eventID := mux.Vars(r)["id"]
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.LogError("updating", err)
	}

	err = json.Unmarshal(body, &updateEvent)

	for i, event := range contracts.Events {
		if event.ID != eventID {
			continue
		}

		event.Title = updateEvent.Title
		event.Description = updateEvent.Description
		contracts.Events = append(contracts.Events[:i], event)
		json.NewEncoder(w).Encode(event)
	}
}
