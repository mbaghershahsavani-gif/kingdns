package monitoring

type ClusterHealth struct {
	Cluster string
	Healthy bool
}

func CheckCluster(cluster string) ClusterHealth {
	return ClusterHealth{
		Cluster: cluster,
		Healthy: true,
	}
}
