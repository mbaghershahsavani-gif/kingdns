package monitoring

type IntelligenceMetric struct {
	Name  string
	Value int
}

func CollectIntelligenceMetrics() []IntelligenceMetric {
	return []IntelligenceMetric{}
}
