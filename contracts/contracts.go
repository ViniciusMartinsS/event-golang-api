package contracts

// Event - Event structure to follow
type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// AllEvents - Array list to save all events
type allEvents []Event

// Events - Initial Database Value
var Events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}
