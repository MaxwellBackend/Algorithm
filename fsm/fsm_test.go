package fsm


import (
	"testing"
)

const (
	STATE_NO_1  = "状态一"
	STATE_NO_2  = "状态二"
	STATE_NO_3  = "状态三"
	STATE_NO_4  = "状态四"
	STATE_NO_5  = "状态五"
)

type StateNo1 struct {
	StateBase
}

func NewStateNo1(fsm *FSM, value string) *StateNo1 {
	state := &StateNo1{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateNo1) Update() {
	if s.fsm.input_hour == 10 {
		s.fsm.ChangeState(STATE_NO_2)
	}
}

type StateNo2 struct {
	StateBase
}

func NewStateNo2(fsm *FSM, value string) *StateNo2 {
	state := &StateNo2{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateNo2) Update() {
	if s.fsm.input_hour == 12 {
		s.fsm.ChangeState(STATE_NO_3)
	}
}

type StateNo3 struct {
	StateBase
}

func NewStateNo3(fsm *FSM, value string) *StateNo3 {
	state := &StateNo3{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateNo3) Update() {
	if s.fsm.input_hour == 14 {
		s.fsm.ChangeState(STATE_NO_4)
	}
}

type StateNo4 struct {
	StateBase
}

func NewStateNo4(fsm *FSM, value string) *StateNo4 {
	state := &StateNo4{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateNo4) Update() {
	if s.fsm.input_hour == 19 {
		s.fsm.ChangeState(STATE_NO_5)
	}
}

type StateNo5 struct {
	StateBase
}

func NewStateNo5(fsm *FSM, value string) *StateNo5 {
	state := &StateNo5{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateNo5) Update() {
	if s.fsm.input_hour == 8 {
		s.fsm.ChangeState(STATE_NO_1)
	}
}

func TestStateMachine(t *testing.T) {
	fsm := NewStateMachine()
	fsm.def_state = STATE_NO_1
	fsm.registers[STATE_NO_1] = NewStateNo1(fsm, STATE_NO_1)
	fsm.registers[STATE_NO_2] = NewStateNo2(fsm, STATE_NO_2)
	fsm.registers[STATE_NO_3] = NewStateNo3(fsm, STATE_NO_3)
	fsm.registers[STATE_NO_4] = NewStateNo4(fsm, STATE_NO_4)
	fsm.registers[STATE_NO_5] = NewStateNo5(fsm, STATE_NO_5)

	fsm.Process(8)
	fsm.Process(10)
	fsm.Process(12)
	fsm.Process(14)
	fsm.Process(19)
	fsm.Process(8)

}
