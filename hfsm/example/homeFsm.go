package main

import "github.com/MaxwellBackend/Algorithm/hfsm"

// 在家状态机
type HomeFsm struct {
	hfsm.FsmBase
}

// 睡觉
type SleepState struct {
	hfsm.StateBase
}

// 玩手机
type PlayPhoneState struct {
	hfsm.StateBase
}

func init() {
	homeFsm := &HomeFsm{}
	homeFsm.Init("HomeFsm", root)

	sleepState := &SleepState{}
	sleepState.Init("SleepState", homeFsm)

	playPhoneState := &PlayPhoneState{}
	playPhoneState.Init("PlayPhoneState", homeFsm)
}