package fsm


import (
	"fmt"
	"testing"
	"time"
)

const (
	STATE_START = "state_start"
	STATE_NO_1  = "state_no_1"
	STATE_NO_2  = "state_no_2"
	STATE_NO_3  = "state_no_3"
)

type StateStart struct {
	StateBase
}

func NewStateStart(fsm *FSM, value string) *StateStart {
	state := &StateStart{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateStart) Start() {
	fmt.Printf("开始状态开始\n")
}

func (s *StateStart) Update(t time.Time) {

}

type StateNo1 struct {
	StateBase
}

func NewStateNo1(fsm *FSM, value string) *StateNo1 {
	state := &StateNo1{}
	state.fsm = fsm
	state.value = value

	return state
}

func (s *StateNo1) Start() {
	fmt.Printf("切换挡位1\n")
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

func (s *StateNo2) Start() {
	fmt.Printf("切换挡位2\n")
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

func (s *StateNo3) Start() {
	fmt.Printf("切换挡位3\n")
}

func TestStateMachine(t *testing.T) {
	fsm := NewStateMachine()
	fsm.def_state = STATE_START
	fsm.registers[STATE_START] = NewStateStart(fsm, STATE_START)
	fsm.registers[STATE_NO_1] = NewStateNo1(fsm, STATE_NO_1)
	fsm.registers[STATE_NO_2] = NewStateNo2(fsm, STATE_NO_2)
	fsm.registers[STATE_NO_3] = NewStateNo3(fsm, STATE_NO_3)
	fsm.Process(time.Now())
	fsm.ChangeState(STATE_NO_1)
	fsm.ChangeState(STATE_NO_2)
	fsm.ChangeState(STATE_NO_3)

}
