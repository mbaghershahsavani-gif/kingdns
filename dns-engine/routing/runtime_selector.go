package routing

type RuntimeNode struct {
	IP      string
	Healthy bool
	Latency int
}

func SelectRuntime(nodes []RuntimeNode) string {
	for _, node := range nodes {
		if node.Healthy {
			return node.IP
		}
	}

	return ""
}
