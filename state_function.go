package state_machine

// StateFunction type defines the function signature that state machine can operate on.
// A state function does not takes any input and returns a state function that is next state.
// An implementing type should implement all state function according to this function signature
// and register to the state machine.
type StateFunction func() StateFunction
