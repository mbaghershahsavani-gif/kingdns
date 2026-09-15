package monitoring

type HealthMetric struct {
	Node    string
	Healthy bool
	Latency int
}
