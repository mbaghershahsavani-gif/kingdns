package routing

type FailoverNode struct {
	IP      string
	Healthy bool
	Latency int
}

func SelectFailover(nodes []FailoverNode) string {
	best := ""

	for _, node := range nodes {
		if node.Healthy {
			if best == "" || node.Latency < 999999 {
				best = node.IP
			}
		}
	}

	return best
}
