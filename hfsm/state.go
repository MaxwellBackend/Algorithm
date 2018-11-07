package hfsm

type StateId string // 状态ID

type StateEvent string // 状态事件

type IState interface {
	Enter()  // 进入
	Update() // 更新
	Exit()   // 退出
}

type StateBase struct {
	Id  StateId
	Fsm IFsm
}

func (sb *StateBase) Init(id StateId, fsm IFsm, self IState) {
	sb.Id = id
	sb.Fsm = fsm

	fsm.RegisterState(id, self)
}

func (sb *StateBase) Enter() {
}

func (sb *StateBase) Update() {
}

func (sb *StateBase) Exit() {

}
