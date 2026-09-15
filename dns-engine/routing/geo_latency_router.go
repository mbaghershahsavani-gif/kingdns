package routing

type RouteCandidate struct {
	IP      string
	Latency int
	Healthy bool
}

func SelectBest(candidates []RouteCandidate) string {
	for _, c := range candidates {
		if c.Healthy {
			return c.IP
		}
	}

	return ""
}
