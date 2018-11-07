package hfsm

type IFsm interface {
	RegisterEvent(event StateEvent, handler EventHandler)
	EventHandle(event StateEvent)
	RegisterState(id StateId, state IState)
	ChangeState(id StateId)
}

type EventHandler func(event StateEvent)

type FsmBase struct {
	StateBase
	Parent               IFsm
	NowStateId           StateId
	registerState        map[StateId]IState
	registerEventHandler map[StateEvent]EventHandler
}

func (fb *FsmBase) Init(id StateId, parent IFsm, self IState) {
	fb.Id = id
	fb.Parent = parent
	fb.registerState = make(map[StateId]IState)
	fb.registerEventHandler = make(map[StateEvent]EventHandler)

	if parent != nil {
		parent.RegisterState(id, self)
	}
}

func (fb *FsmBase) Update() {
	nowState := fb.registerState[fb.NowStateId]
	nowState.Update()
}

func (fb *FsmBase) Exit() {
	if fb.NowStateId != "" {
		nowState := fb.registerState[fb.NowStateId]
		nowState.Exit()
	}

	fb.ResetState()
}

func (fb *FsmBase) RegisterState(id StateId, state IState) {
	fb.registerState[id] = state
}

func (fb *FsmBase) ChangeState(id StateId) {
	// 未改变状态，直接返回
	if id == fb.NowStateId {
		return
	}

	if fb.NowStateId != "" {
		nowState := fb.registerState[fb.NowStateId]
		nowState.Exit()
	}

	state := fb.registerState[id]
	fb.NowStateId = id

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
	fb.NowStateId = ""
}
