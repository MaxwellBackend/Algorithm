package hfsm

type StateId string

type StateEvent string

type IState interface {
	Enter()
	Update()
	Exit()
}

type StateBase struct {
	Id StateId
	Fsm IFsm
}

func (sb *StateBase) Init(id StateId, fsm IFsm) {
	sb.Id = id
	sb.Fsm = fsm

	fsm.RegisterState(id, sb)
}

func (sb *StateBase) Enter() {

}

func (sb *StateBase) Update() {

}

func (sb *StateBase) Exit() {

}