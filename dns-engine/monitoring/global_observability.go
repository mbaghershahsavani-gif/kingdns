package monitoring

type GlobalMetric struct {
	Region string
	Value  int
}

func CollectGlobalMetrics() []GlobalMetric {
	return []GlobalMetric{}
}
