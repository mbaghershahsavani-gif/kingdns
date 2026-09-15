package monitoring

type ProbeResult struct {
	Node    string
	Latency int
	Healthy bool
}

func Probe(node string) ProbeResult {
	return ProbeResult{
		Node:    node,
		Latency: 1,
		Healthy: true,
	}
}
