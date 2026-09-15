package monitoring

type DashboardMetric struct {
	Queries uint64
	Errors  uint64
	Latency uint64
}

func Export(metric DashboardMetric) {
	// Future:
	// expose Prometheus metrics
}
