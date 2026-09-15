package routing

type Candidate struct {
	NodeID uint
	Latency int
	Health float64
}

func Select(candidates []Candidate) uint {
	if len(candidates) == 0 {
		return 0
	}

	best := candidates[0]

	for _, node := range candidates {
		if node.Health > best.Health ||
			(node.Health == best.Health && node.Latency < best.Latency) {
			best = node
		}
	}

	return best.NodeID
}
