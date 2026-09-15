package agent

type Heartbeat struct {
	NodeID int
	Health int
	Latency int
}

func Send(h Heartbeat) {
	// Future:
	// send agent health to controller
}
