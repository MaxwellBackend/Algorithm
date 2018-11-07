package main

import (
	"github.com/MaxwellBackend/Algorithm/hfsm"
	"time"
	"fmt"
)

type Root struct {
	hfsm.FsmBase
}

var root *Root

var now time.Time

func init () {
	root = &Root{}
	root.Init("Parent", nil, root)

	homeInit()
	travelInit()
	companyInit()
}

func log(msg string) {
	fmt.Printf("[%v] %v\n", now, msg)
}

func main() {
	now = time.Date(2018,11,6,6,0,0,0, time.Local)
	root.ChangeState("HomeFsm")
	for {
		root.Update()
		time.Sleep(time.Second / 20)
		now = now.Add(1 * time.Minute)
	}
}