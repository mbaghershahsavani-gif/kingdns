package monitoring

type QueryMetric struct {
	Domain  string
	Latency int
	Status  string
}

func Publish(metric QueryMetric) {
	// Future:
	// send metrics to controller
}
