package fsm


import (
"fmt"
)

// 状态接口
type IState interface {
	Start()
	Update()
	Stop()
	Value() string
}

// 状态基类
type StateBase struct {
	fsm   *FSM
	value string
}

func (b *StateBase) Start() {
	fmt.Printf("%v开始\n",b.Value())
}

func (b *StateBase) Update() {

}

func (b *StateBase) Stop() {
	fmt.Printf("%v结束\n", b.value)
}

func (b *StateBase) Value() string {
	return b.value
}

type FSM struct {
	state     IState            // 当前状态
	registers map[string]IState // 状态集合
	def_state string            // 状态集合
	input_hour uint32 			// 输入的时间
}

func NewStateMachine() *FSM {
	fsm := &FSM{}
	fsm.registers = make(map[string]IState)
	return fsm
}

// 执行
func (fsm *FSM) Process(hour uint32)  error {
	fsm.input_hour = hour
	if fsm.state == nil {
		fsm.ChangeState(fsm.def_state)
	} else {
		fsm.state.Update()
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

