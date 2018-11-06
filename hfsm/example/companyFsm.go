package main

import (
	"github.com/MaxwellBackend/Algorithm/hfsm"
	"math/rand"
)

const EventCodeEnd = "EventCodeEnd"
const EventReadEnd = "EventReadEnd"
const EventMeetEnd = "EventMeetEnd"

// 在公司状态机
type CompanyFsm struct {
	hfsm.FsmBase
}

func (f *CompanyFsm) Enter(id hfsm.StateId) {
	if id == "" {
		id = "CodeState"
	}

	f.FsmBase.Enter(id)
}

func (f *CompanyFsm) Init(id hfsm.FsmId, root hfsm.IRoot, self hfsm.IFsm) {
	f.FsmBase.Init(id, root, self)

	f.RegisterEvent(EventCodeEnd, f.handleRandom)
	f.RegisterEvent(EventReadEnd, f.handleRandom)
	f.RegisterEvent(EventMeetEnd, f.handleRandom)
}

func (f *CompanyFsm) Update() {
	if now.Hour() >= 18 && now.Minute() >= 30 {
		f.ResetState()
		log("下班啦")
		f.Root.ChangeFsm("TravelFsm", "S2CWalkState")
		return
	}

	f.FsmBase.Update()
}

func (f *CompanyFsm) handleRandom(event hfsm.StateEvent) {
	var states = []hfsm.StateId{"CodeState", "ReadState", "MeetState"}
	index := rand.Uint32() % uint32(len(states))

	f.ChangeState(states[index])
}

// 敲代码状态
type CodeState struct {
	hfsm.StateBase
	startTime int
}

func (s *CodeState) Enter() {
	s.startTime = 0
	log("开始敲代码")
}

func (s *CodeState) Update() {
	s.startTime++
	if s.startTime >= 60 {
		s.Fsm.EventHandle(EventCodeEnd)
	}
}

func (s *CodeState) Exit() {
	log("结束敲代码")
}

// 读书
type ReadState struct {
	hfsm.StateBase
	startTime int
}

func (s *ReadState) Enter() {
	s.startTime = 0
	log("开始读书")
}

func (s *ReadState) Update() {
	s.startTime++
	if s.startTime >= 30 {
		s.Fsm.EventHandle(EventReadEnd)
	}
}

func (s *ReadState) Exit() {
	log("结束读书")
}

// 开会状态
type MeetState struct {
	hfsm.StateBase
	startTime int
}

func (s *MeetState) Enter() {
	s.startTime = 0
	log("开始开会")
}

func (s *MeetState) Update() {
	s.startTime++
	if s.startTime >= 60 {
		s.Fsm.EventHandle(EventMeetEnd)
	}
}

func (s *MeetState) Exit() {
	log("结束开会")
}

func companyInit() {
	companyFsm := &CompanyFsm{}
	companyFsm.Init("CompanyFsm", root, companyFsm)

	codeState := &CodeState{}
	codeState.Init("CodeState", companyFsm, codeState)

	readState := &ReadState{}
	readState.Init("ReadState", companyFsm, readState)

	meetState := &MeetState{}
	meetState.Init("MeetState", companyFsm, meetState)
}