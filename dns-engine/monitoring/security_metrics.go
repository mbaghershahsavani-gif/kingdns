package monitoring

type SecurityMetric struct {
	Name  string
	Value int
}

func CollectSecurityMetrics() []SecurityMetric {
	return []SecurityMetric{}
}
