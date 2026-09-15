package database

type NodeRecord struct {
	Name    string
	IP      string
	Healthy bool
	Latency int
}

func HealthyNodes() []NodeRecord {
	return []NodeRecord{
		{
			Name:    "local-node",
			IP:      "127.0.0.1",
			Healthy: true,
			Latency: 1,
		},
	}
}
