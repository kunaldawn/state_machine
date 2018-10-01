package state_machine

import (
	"encoding/json"
	"errors"
)

// StateMachine handles execution of state handlers and also automatic persistence and recovery of the state handler.
type StateMachine struct {
	stateHandler       StateHandler            // state handler that the state machine need to operate on
	persistenceManager StatePersistenceManager // persistence manager that the state machine uses to persist the state and recover the state
	stateMap           map[string]*State       //  state map stores all the registered states by the state handler, its generally used to trigger the persisted state
	currentState       *State                  // current state denotes the current state the state manager is executing, its self forwarding
}

// NewStateMachine gives an instance of state machine given a state handler and persistence manager
func NewStateMachine(stateHandler StateHandler, persistenceManager StatePersistenceManager) (*StateMachine, error) {
	// create a state machine and initialize it
	machine := &StateMachine{stateHandler: stateHandler, persistenceManager: persistenceManager}
	err := machine.init()

	return machine, err
}

func (machine *StateMachine) init() error {
	// get starting state from the state handler and set it as machines current state
	machine.currentState = machine.stateHandler.GetStartingState()

	// create a new state map and register all states given by the state handler
	machine.stateMap = make(map[string]*State)
	// get all states from state handler to register
	states := machine.stateHandler.GetStates()
	// store all registered states to state map
	for _, state := range states {
		machine.stateMap[state.Name()] = state
	}

	// check if valid persistence manager exist, then try to restore the state
	if machine.persistenceManager != nil {
		// try to load the persisted state
		data, err := machine.persistenceManager.Load()

		// check if any error is there, then return the error
		if err != nil {
			return err
		}

		// check if any data is valid try to load it
		if data != nil && len(data) > 0 {
			// create a new state data and load the state from persisted data
			machineState := &StateData{}
			err := json.Unmarshal(data, machineState)

			// check if there is any loading error
			if err != nil {
				return err
			}

			// generate the handler state data from machine state data
			handlerData, err := json.Marshal(machineState.StateData)

			// check if there is an error while generating data for state handler
			if err != nil {
				return err
			}

			// ask state handler to load the state
			err = machine.stateHandler.UnmarshalJSON(handlerData)

			// check if there is any error generated by state handler
			if err != nil {
				return err
			}

			// check if last persisted state is registered
			if currentState, ok := machine.stateMap[machineState.StateName]; ok {
				// set machines current
				machine.currentState = currentState
			} else {
				// persisted state is not registered by state handler
				return errors.New("starting state is not registered")
			}
		}
	}

	return nil
}

// Run runs the state machine and state cycle is started or resumed at the state handler.
// state handler can mark nil state to finish execution of this blocking call.
// state handler data and state machine state is persisted before each state execution.
func (machine *StateMachine) Run(isDelete bool) error {
	// if machine has a current valid state
	for machine.currentState != nil && machine.currentState.Valid() {
		// create a new state data to persist
		stateData := StateData{StateData: machine.stateHandler, StateName: machine.currentState.Name()}
		// generate data for persistence manager
		data, err := json.Marshal(stateData)

		// check if there is any error in creating data for state handler
		if err != nil {
			return err
		}

		// ask persistence manager to save the state
		err = machine.persistenceManager.Save(data)

		// check if persistence manager has reported any errors
		if err != nil {
			return err
		}

		// execute machines current state
		machine.currentState.Do()
	}

	// all execution has finished, check if state delete is required
	if isDelete {
		// delete the machines state as its not required any more
		machine.persistenceManager.Delete()
	}

	return nil
}