package monitoring

type TenantMetric struct {
	TenantID string
	Queries  int
}

func CollectTenantMetrics(id string) TenantMetric {
	return TenantMetric{
		TenantID: id,
	}
}
