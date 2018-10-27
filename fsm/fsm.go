package fsm


import (
"fmt"
"time"
)

// 状态接口
type IState interface {
	Start()
	Update(time.Time)
	Stop()
	Value() string
}

// 状态基类
type StateBase struct {
	fsm   *FSM
	value string
}

func (b *StateBase) Start() {
	fmt.Printf("Start %v\n", b.value)
}

func (b *StateBase) Update(t time.Time) {

}

func (b *StateBase) Stop() {
	fmt.Printf("Stop %v\n", b.value)
}

func (b *StateBase) Value() string {
	return b.value
}

type FSM struct {
	state     IState            // 当前状态
	registers map[string]IState // 状态集合
	def_state string            // 状态集合
}

func NewStateMachine() *FSM {
	fsm := &FSM{}
	fsm.registers = make(map[string]IState)
	return fsm
}

// 执行
func (fsm *FSM) Process(t time.Time) error {
	if fsm.state == nil {
		fsm.ChangeState(fsm.def_state)
		fmt.Printf("FSM: init state %v\n", fsm.state.Value())
	} else {
		fsm.state.Update(t)
	}

	return nil
}

// 切换状态
func (fsm *FSM) ChangeState(value string) {
	if fsm.state == nil {
		fsm.state = fsm.registers[value]
	} else if fsm.state.Value() != value {
		fsm.state.Stop()
		fsm.state = fsm.registers[value]
	} else {
		fmt.Printf("同状态不用切换\n")
		return
	}

	fsm.state.Start()
}

