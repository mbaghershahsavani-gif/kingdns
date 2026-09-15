package routing

type Target struct {
	Name string
	Latency int
	Health int
}

func Select(targets []Target) Target {
	if len(targets) == 0 {
		return Target{}
	}

	return targets[0]
}
