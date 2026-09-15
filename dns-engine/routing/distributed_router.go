package routing

type DistributedNode struct {
	IP      string
	Healthy bool
	Latency int
}

func SelectDistributed(nodes []DistributedNode) string {
	for _, node := range nodes {
		if node.Healthy {
			return node.IP
		}
	}

	return ""
}
