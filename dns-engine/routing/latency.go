package routing

type Candidate struct {
	Name string
	Latency int
	Health int
}

func Best(candidates []Candidate) Candidate {
	if len(candidates) == 0 {
		return Candidate{}
	}

	best := candidates[0]

	for _, node := range candidates {
		if node.Latency < best.Latency && node.Health > 0 {
			best = node
		}
	}

	return best
}
