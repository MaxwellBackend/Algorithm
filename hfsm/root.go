package hfsm

type IRoot interface {
	Init()
	Update()
	ChangeFsm(id FsmId)
	RegisterFsm(id FsmId, fsm IFsm)
}

type Root struct {
	nowFsmId FsmId
	registerFsm map[FsmId]IFsm
}

func (r *Root) Init() {
	r.registerFsm = make(map[FsmId]IFsm)
}

func (r *Root) Update() {
	nowFsm := r.registerFsm[r.nowFsmId]
	nowFsm.Update()
}

func (r *Root) RegisterFsm(id FsmId, fsm IFsm) {
	r.registerFsm[id] = fsm
}

func (r *Root) ChangeFsm(id FsmId) {
	fsm := r.registerFsm[id]
	r.nowFsmId = id
	fsm.Enter()
}