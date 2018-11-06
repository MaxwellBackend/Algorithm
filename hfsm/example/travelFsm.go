package main

import "github.com/MaxwellBackend/Algorithm/hfsm"

// 在路上状态
type TravelFsm struct {
	hfsm.FsmBase
}

// 地铁到公司的走路状态
type S2CWalkState struct {
	hfsm.StateBase
}

// 地铁到家的走路状态
type S2HWalkState struct {
	hfsm.StateBase
}

// 坐地铁状态
type SubwayState struct {
	hfsm.StateBase
}

func init() {
	travelFsm := &TravelFsm{}
	travelFsm.Init("TravelFsm", root)

	s2cWalkState := &S2CWalkState{}
	s2cWalkState.Init("S2CWalkState", travelFsm)

	s2hWalkState := &S2HWalkState{}
	s2hWalkState.Init("S2HWalkState", travelFsm)

	subwayState := &SubwayState{}
	subwayState.Init("SubwayState", travelFsm)
}