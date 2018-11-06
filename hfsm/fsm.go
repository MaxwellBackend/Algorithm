package hfsm

type FsmId string

type IFsm interface {
	Enter()
	Update()
	Exit()
	RegisterEvent(event StateEvent, handler EventHandler)
	EventHandle(event StateEvent)
	RegisterState(id StateId, state IState)
	ChangeState(id StateId)
}

type EventHandler func(event StateEvent)

type FsmBase struct {
	Id                   FsmId
	Root                 IRoot
	nowStateId           StateId
	registerState        map[StateId]IState
	registerEventHandler map[StateEvent]EventHandler
}

func (fb *FsmBase) Init(id FsmId, root IRoot) {
	fb.Id = id
	fb.Root = root
	fb.registerState = make(map[StateId]IState)
	fb.registerEventHandler = make(map[StateEvent]EventHandler)

	root.RegisterFsm(id, fb)
}

func (fb *FsmBase) Enter() {

}

func (fb *FsmBase) Update() {
	nowState := fb.registerState[fb.nowStateId]
	nowState.Update()
}

func (fb *FsmBase) Exit() {

}

func (fb *FsmBase) RegisterState(id StateId, state IState) {
	fb.registerState[id] = state
}

func (fb *FsmBase) ChangeState(id StateId) {
	state := fb.registerState[id]
	fb.nowStateId = id
	state.Enter()
}

func (fb *FsmBase) RegisterEvent(event StateEvent, handler EventHandler) {
	fb.registerEventHandler[event] = handler
}

func (fb *FsmBase) EventHandle(event StateEvent) {
	handler := fb.registerEventHandler[event]
	handler(event)
}