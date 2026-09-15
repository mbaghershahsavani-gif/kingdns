package intelligence

type PerformanceProfile struct {
	Node    string
	Latency int
}

func Profile(node string) PerformanceProfile {
	return PerformanceProfile{
		Node: node,
	}
}
