package main

import (
	"github.com/MaxwellBackend/Algorithm/hfsm"
)

const EventArriveCompany = "ArriveCompany"
const EventArriveHome = "ArriveHome"

// 在路上状态
type TravelFsm struct {
	hfsm.FsmBase
	oldState hfsm.StateId
}

func (f *TravelFsm) Init(id hfsm.StateId, parent hfsm.IFsm, self hfsm.IState) {
	f.FsmBase.Init(id, parent, self)

	f.RegisterEvent(EventArriveCompany, f.handleEventArriveCompany)
	f.RegisterEvent(EventArriveHome, f.handleEventArriveHome)
}

func (f *TravelFsm) Enter() {
	if f.oldState == "" {
		f.ChangeState("S2HWalkState")
	} else {
		f.ChangeState(f.oldState)
	}
}

func (f *TravelFsm) Exit() {
	f.oldState = f.NowStateId
	f.FsmBase.Exit()
}

func (f *TravelFsm) handleEventArriveCompany(event hfsm.StateEvent) {
	f.Parent.ChangeState("CompanyFsm")
}

func (f *TravelFsm) handleEventArriveHome(event hfsm.StateEvent) {
	f.Parent.ChangeState("HomeFsm")
}

// 地铁到公司的走路状态
type S2CWalkState struct {
	hfsm.StateBase
	startTime int
}

func (s *S2CWalkState) Enter() {
	s.startTime = 0
	if now.Hour() < 12 {
		log("从桂林路站走到公司")
	} else {
		log("从公司走到桂林路站")
	}
}

func (s *S2CWalkState) Update() {
	s.startTime++
	if s.startTime >= 10 {
		if now.Hour() < 12 {
			s.Fsm.EventHandle(EventArriveCompany)
		} else {
			s.Fsm.ChangeState("SubwayState")
		}
	}
}

func (s *S2CWalkState) Exit() {
	if now.Hour() < 12 {
		log("到达公司")
	} else {
		log("到达桂林路站")
	}
}

// 地铁到家的走路状态
type S2HWalkState struct {
	hfsm.StateBase
	startTime int
}

func (s *S2HWalkState) Enter() {
	s.startTime = 0
	if now.Hour() < 12 {
		log("从家走到地铁")
	} else {
		log("从地铁走到家")
	}
}

func (s *S2HWalkState) Update() {
	s.startTime++
	if s.startTime >= 20 {
		if now.Hour() < 12 {
			s.Fsm.ChangeState("SubwayState")
		} else {
			s.Fsm.EventHandle(EventArriveHome)
		}
	}
}

func (s *S2HWalkState) Exit() {
	if now.Hour() < 12 {
		log("到达松江站")
	} else {
		log("到家")
	}
}

// 坐地铁状态
type SubwayState struct {
	hfsm.StateBase
	startTime int
}

func (s *SubwayState) Enter() {
	s.startTime = 0
	if now.Hour() < 12 {
		log("地铁：从松江站到桂林路站")
	} else {
		log("地铁：从桂林路站到松江站")
	}
}

func (s *SubwayState) Update() {
	s.startTime++
	if s.startTime >= 30 {
		if now.Hour() < 12 {
			s.Fsm.ChangeState("S2CWalkState")
		} else {
			s.Fsm.ChangeState("S2HWalkState")
		}
	}
}

func (s *SubwayState) Exit() {
	log("走出地铁")
}

func travelInit() {
	travelFsm := &TravelFsm{}
	travelFsm.Init("TravelFsm", root, travelFsm)

	s2cWalkState := &S2CWalkState{}
	s2cWalkState.Init("S2CWalkState", travelFsm, s2cWalkState)

	s2hWalkState := &S2HWalkState{}
	s2hWalkState.Init("S2HWalkState", travelFsm,s2hWalkState)

	subwayState := &SubwayState{}
	subwayState.Init("SubwayState", travelFsm, subwayState)
}
