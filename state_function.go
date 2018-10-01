/*
  ____  ___    ____  _        _         __  __            _     _
 / ___|/ _ \  / ___|| |_ __ _| |_ ___  |  \/  | __ _  ___| |__ (_)_ __   ___
| |  _| | | | \___ \| __/ _` | __/ _ \ | |\/| |/ _` |/ __| '_ \| | '_ \ / _ \
| |_| | |_| |  ___) | || (_| | ||  __/ | |  | | (_| | (__| | | | | | | |  __/
 \____|\___/  |____/ \__\__,_|\__\___| |_|  |_|\__,_|\___|_| |_|_|_| |_|\___|

*/

package state_machine

// StateFunction type defines the function signature that state machine can operate on.
// A state function does not takes any input and returns a state function that is next state.
// An implementing type should implement all state function according to this function signature
// and register to the state machine.
type StateFunction func() StateFunction
