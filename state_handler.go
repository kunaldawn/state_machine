package state_machine

import "encoding/json"

// StateHandler defines the interface that a type need to implement in order to be handled by state machine.
type StateHandler interface {
	json.Marshaler   // state handlers need to implement Marshaler interface so that state data can be persisted by the state machine
	json.Unmarshaler // state handlers need to implement Unmarshaler interface so that persisted state data can be loaded back by state machine

	// GateStates interface need to be implemented by state handlers so that it can register all state functions to
	// the state machine. It should return a slice of State pointers which need to be registered to the state machine.
	// State handlers should register all state functions so that state machine can resume state.
	GetStates() []*State

	// GetStartingState interface need to be implemented by state handlers so that state machine can know
	// about the starting state of the state handler. In case of state handler is recovered from persistence
	// manager, starting state of the recovered state handler is determined by the state machine itself.
	GetStartingState() *State
}
