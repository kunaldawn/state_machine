package state_machine

import (
	"reflect"
	"runtime"
	"strings"
)

// State defines state that state machine can operate on
type State struct {
	stateFunction StateFunction // state function for the state
}

// NewState gives a state given a state function to operate on
func NewState(stateFunction StateFunction) *State {
	return &State{stateFunction: stateFunction}
}

// Name gives name of the state, its generally not exactly same as function name but rather
// internal implementation of the state function
func (state *State) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(state.stateFunction).Pointer()).Name()
	parts := strings.Split(name, "/")

	return parts[len(parts)-1]
}

// Do function executes the state function and forwards the state function to next state
// If the state is nil, then no execution is done
func (state *State) Do() {
	// check that state function is not nil
	if state.stateFunction != nil {
		// call the state function and get next state function
		stateFunction := state.stateFunction()
		// replace current state function with next state function
		state.stateFunction = stateFunction
	}
}

// Valid returns boolean flag indicating true if state is executable, false otherwise
func (state *State) Valid() bool {
	return state.stateFunction != nil
}
