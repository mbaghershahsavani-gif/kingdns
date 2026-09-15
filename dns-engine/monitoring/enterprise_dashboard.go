package monitoring

type DashboardMetric struct {
	Name  string
	Value int
}

func CollectDashboard() []DashboardMetric {
	return []DashboardMetric{}
}
