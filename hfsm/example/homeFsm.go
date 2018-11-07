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

func (f *HomeFsm) Enter() {
	if now.Hour() > 22 || now.Hour() < 8 {
		f.ChangeState("SleepState")
	} else {
		f.ChangeState("PlayPhoneState")
	}
}

func (f *HomeFsm) Init(id hfsm.StateId, parent hfsm.IFsm, self hfsm.IState) {
	f.FsmBase.Init(id, parent, self)

	f.RegisterEvent(EventWashEnd, f.handleEventWashEnd)
}

func (f *HomeFsm) handleEventWashEnd(event hfsm.StateEvent) {
	f.Parent.ChangeState("TravelFsm")
}

// 睡觉
type SleepState struct {
	hfsm.StateBase
}

func (s *SleepState) Enter() {
	log("睡觉")
}

func (s *SleepState) Update() {
	if now.Hour() == 8 && now.Minute() == 0 {
		s.Fsm.ChangeState("WashState")
	}
}

func (s *SleepState) Exit() {
	log("起床")
}

// 玩手机
type PlayPhoneState struct {
	hfsm.StateBase
	startTime int
}

func (s *PlayPhoneState) Enter() {
	s.startTime = 0
	log("开始玩手机")
}

func (s *PlayPhoneState) Update() {
	s.startTime++
	if s.startTime >= 120 {
		s.Fsm.ChangeState("SleepState")
	}
}

func (s *PlayPhoneState) Exit() {
	log("不玩手机了")
}

// 洗漱
type WashState struct {
	hfsm.StateBase
	startTime int
}

func (s *WashState) Enter() {
	s.startTime = 0
	log("洗漱")
}

func (s *WashState) Update() {
	s.startTime++
	if s.startTime >= 30 {
		s.Fsm.EventHandle(EventWashEnd)
	}
}

func (s *WashState) Exit() {
	log("出门")
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
