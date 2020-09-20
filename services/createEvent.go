package services

import (
	"encoding/json"
	"fmt"
	"go-crud-api/contracts"
	"go-crud-api/utils"
	"io/ioutil"
	"net/http"
)

// CreateEvent - Responsible for creating an event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent contracts.Event
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.LogError("creating", err)
	}

	err = json.Unmarshal(body, &newEvent)

	if err != nil {
		utils.LogError("creating", err)
	}

	err = checkExistingID(newEvent.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID already exists")
		return
	}

	contracts.Events = append(contracts.Events, newEvent)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func checkExistingID(ID string) error {
	var err error
	allEvents := contracts.Events

	for _, event := range allEvents {
		if event.ID == ID {
			err = fmt.Errorf("ID already exists %s", ID)
			break
		}

		err = nil
	}

	return err
}
