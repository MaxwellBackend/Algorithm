package hfsm

type FsmId string

type IFsm interface {
	Enter(id StateId)
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

func (fb *FsmBase) Init(id FsmId, root IRoot, self IFsm) {
	fb.Id = id
	fb.Root = root
	fb.registerState = make(map[StateId]IState)
	fb.registerEventHandler = make(map[StateEvent]EventHandler)

	root.RegisterFsm(id, self)
}

func (fb *FsmBase) Enter(id StateId) {
	fb.ChangeState(id)
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
	// 未改变状态，直接返回
	if id == fb.nowStateId {
		return
	}

	if fb.nowStateId != "" {
		nowState := fb.registerState[fb.nowStateId]
		nowState.Exit()
	}

	state := fb.registerState[id]
	fb.nowStateId = id

	state.Enter()
}

func (fb *FsmBase) RegisterEvent(event StateEvent, handler EventHandler) {
	fb.registerEventHandler[event] = handler
}

func (fb *FsmBase) EventHandle(event StateEvent) {
	handler, found := fb.registerEventHandler[event]
	if found {
		handler(event)
	}
}

func (fb *FsmBase) StateCount() int {
	return len(fb.registerState)
}

func (fb *FsmBase) ResetState() {
	fb.nowStateId = ""
}