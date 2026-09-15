package routing

type RegionNode struct {
	Region  string
	IP      string
	Latency int
	Healthy bool
}

func SelectRegion(nodes []RegionNode) string {
	var selected string
	best := 999999

	for _, node := range nodes {
		if node.Healthy && node.Latency < best {
			best = node.Latency
			selected = node.IP
		}
	}

	return selected
}
