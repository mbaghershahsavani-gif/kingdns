package monitoring

type ProbeScheduler struct {
	Interval int
}

func (p ProbeScheduler) Run() string {
	return "probe-complete"
}
