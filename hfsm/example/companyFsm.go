package main

import "github.com/MaxwellBackend/Algorithm/hfsm"

// 在公司状态机
type CompanyFsm struct {
	hfsm.FsmBase
}

// 工作状态
type WorkState struct {
	hfsm.StateBase
}

// 读书
type ReadState struct {
	hfsm.StateBase
}

// 开会状态
type MeetState struct {
	hfsm.StateBase
}

func init() {
	companyFsm := &CompanyFsm{}
	companyFsm.Init("CompanyFsm", root)

	workState := &WorkState{}
	workState.Init("WorkState", companyFsm)

	readState := &ReadState{}
	readState.Init("ReadState", companyFsm)

	meetState := &MeetState{}
	meetState.Init("MeetState", companyFsm)
}