/*
  ____  ___    ____  _        _         __  __            _     _
 / ___|/ _ \  / ___|| |_ __ _| |_ ___  |  \/  | __ _  ___| |__ (_)_ __   ___
| |  _| | | | \___ \| __/ _` | __/ _ \ | |\/| |/ _` |/ __| '_ \| | '_ \ / _ \
| |_| | |_| |  ___) | || (_| | ||  __/ | |  | | (_| | (__| | | | | | | |  __/
 \____|\___/  |____/ \__\__,_|\__\___| |_|  |_|\__,_|\___|_| |_|_|_| |_|\___|

*/

package state_machine

// StateData is used by state machine for persistence purpose
type StateData struct {
	StateName string      `json:"state_name"` // name of the state at the time of persistence is triggered
	StateData interface{} `json:"state_data"` // state data that need to be persisted
}
