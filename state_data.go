package state_machine

// StateData is used by state machine for persistence purpose
type StateData struct {
	StateName string      `json:"state_name"` // name of the state at the time of persistence is triggered
	StateData interface{} `json:"state_data"` // state data that need to be persisted
}
