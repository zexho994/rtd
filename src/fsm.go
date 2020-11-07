package src

import (
	"fmt"
	"sync"
)

// RTDState list
const (
	Listen      = RTDState("listen")
	SynSend     = RTDState("syn_send")
	SynRecv     = RTDState("syn_recv")
	Established = RTDState("established")
	FinWait1    = RTDState("fin_wait_1")
	FinWait2    = RTDState("fin_wait_2")
	CloseWait   = RTDState("close_wait")
	LastAck     = RTDState("last_ack")
	TimeWait    = RTDState("time_wait")
	Closed      = RTDState("closed")
)

// RTDEvent list
const (
	// C: Closed -> SynSend
	// S: Closed -> Listen
	SendSyn = RTDEvent("send syn")

	// S: Listen -> SynRecv
	SendSynAck = RTDEvent("send syn ack")

	// C: SynSend -> Established
	// S: CloseWait -> LastAck
	// C: FinWait2 -> TimeWait
	// S: LastAck -> Closed
	SendAck = RTDEvent("send ack")

	// C: Established -> FinWait1
	SendFin = RTDEvent("send fin")

	// S: Established -> CloseWait
	SendFinAck = RTDEvent("send fin ack")

	// C: TimeWait -> Closed
	NotAck = RTDEvent("not receive server ack")
)

type RTDState string            // 状态
type RTDEvent string            // 事件
type RTDHandler func() RTDState // 处理方法

type RTDFsm struct {
	mutex    sync.Mutex                           // 排它锁
	state    RTDState                             // 当前状态
	handlers map[RTDState]map[RTDEvent]RTDHandler // 处理地图集
}

// get state
func (r *RTDFsm) getState() RTDState {
	return r.state
}

// set state
func (r *RTDFsm) setState(state RTDState) {
	r.state = state
}

// handle events
func (r *RTDFsm) Call(event RTDEvent) RTDState {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	events := r.handlers[r.getState()]
	if events == nil {
		return r.getState()
	}

	if fn, ok := events[event]; ok {
		oldState := r.getState()
		r.setState(fn())
		newState := r.getState()
		fmt.Printf("fms state chagned, [%v] -> [%v] \n", oldState, newState)
	}

	return r.getState()
}

// constructor for RTDFsm
func NewRTDFsm(initState RTDState) *RTDFsm {
	return &RTDFsm{
		state:    initState,
		handlers: make(map[RTDState]map[RTDEvent]RTDHandler),
	}
}
