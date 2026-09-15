package controller

type HeartbeatScheduler struct {
	Node string
}

func (h *HeartbeatScheduler) Send() string {
	return "healthy"
}
