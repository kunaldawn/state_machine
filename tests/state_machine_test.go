/*
  ____  ___    ____  _        _         __  __            _     _
 / ___|/ _ \  / ___|| |_ __ _| |_ ___  |  \/  | __ _  ___| |__ (_)_ __   ___
| |  _| | | | \___ \| __/ _` | __/ _ \ | |\/| |/ _` |/ __| '_ \| | '_ \ / _ \
| |_| | |_| |  ___) | || (_| | ||  __/ | |  | | (_| | (__| | | | | | | |  __/
 \____|\___/  |____/ \__\__,_|\__\___| |_|  |_|\__,_|\___|_| |_|_|_| |_|\___|

*/

package tests

import (
	"github.com/kunaldawn/state_machine"
	"testing"
)
import "encoding/json"

type MockStateData struct {
	Name    string `json:"name"`
	Counter int64  `json:"counter"`
}

type MockStateHandler struct {
	state_machine.StateHandler

	handlerData *MockStateData
}

func NewMockHandler(name string) *MockStateHandler {
	return &MockStateHandler{handlerData: &MockStateData{Name: name}}
}

func (handler *MockStateHandler) StateA() state_machine.StateFunction {
	handler.handlerData.Counter++

	return handler.StateB
}

func (handler *MockStateHandler) StateB() state_machine.StateFunction {
	handler.handlerData.Counter++

	return nil
}

func (handler *MockStateHandler) GetStartingState() *state_machine.State {
	return state_machine.NewState(handler.StateA)
}

func (handler *MockStateHandler) GetStates() []*state_machine.State {
	return []*state_machine.State{
		state_machine.NewState(handler.StateA),
		state_machine.NewState(handler.StateB),
	}
}

func (handler *MockStateHandler) MarshalJSON() ([]byte, error) {
	return json.Marshal(handler.handlerData)
}

func (handler *MockStateHandler) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, handler.handlerData)
}

func Test_StateMachine(test *testing.T) {
	handler := NewMockHandler("mock_test")
	stateMachine, err := state_machine.NewStateMachine(handler, nil)

	if err != nil {
		test.Error(err)
	}

	stateMachine.Run(false)

	if handler.handlerData.Counter != 2 || handler.handlerData.Name != "mock_test" {
		test.Errorf("invalid state : %d, %s", handler.handlerData.Counter, handler.handlerData.Name)
	}
}
