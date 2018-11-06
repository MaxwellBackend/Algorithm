package main

import (
	"github.com/MaxwellBackend/Algorithm/hfsm"
)

const (
	EventWashEnd hfsm.StateEvent = "WashEnd"
)

// 在家状态机
type HomeFsm struct {
	hfsm.FsmBase
}

func (f *HomeFsm) Enter(id hfsm.StateId) {
	if id == "" {
		id = "SleepState"
	}

	f.FsmBase.Enter(id)
}

func (f *HomeFsm) Init(id hfsm.FsmId, root hfsm.IRoot, self hfsm.IFsm) {
	f.FsmBase.Init(id, root, self)
	
	f.RegisterEvent(EventWashEnd, f.handleEventWashEnd)
}

func (f *HomeFsm) handleEventWashEnd(event hfsm.StateEvent) {
	f.ResetState()
	f.Root.ChangeFsm("TravelFsm", "S2HWalkState")
}

// 睡觉
type SleepState struct {
	hfsm.StateBase
}

func (s *SleepState) Enter() {
	log("开始睡觉")
}

func (s *SleepState) Update() {
	if now.Hour() == 7 && now.Minute() == 0 {
		s.Fsm.ChangeState("WashState")
	}
}

func (s *SleepState) Exit() {
	log("起床")
}

// 玩手机
type PlayPhoneState struct {
	hfsm.StateBase
}

// 洗漱
type WashState struct {
	hfsm.StateBase
	startTime int
}

func (s *WashState) Enter() {
	s.startTime = 0
	log("开始洗漱")
}

func (s *WashState) Update() {
	s.startTime++
	if s.startTime >= 20 {
		s.Fsm.EventHandle(EventWashEnd)
	}
}

func (s *WashState) Exit() {
	log("结束洗漱")
}


func homeInit() {
	homeFsm := &HomeFsm{}
	homeFsm.Init("HomeFsm", root, homeFsm)

	sleepState := &SleepState{}
	sleepState.Init("SleepState", homeFsm, sleepState)

	playPhoneState := &PlayPhoneState{}
	playPhoneState.Init("PlayPhoneState", homeFsm, playPhoneState)

	washState := &WashState{}
	washState.Init("WashState", homeFsm, washState)
}