package monitoring

type NodeHealth struct {
	Node    string
	Healthy bool
	Latency int
}

func CheckNode(node string) NodeHealth {
	return NodeHealth{
		Node:    node,
		Healthy: true,
		Latency: 0,
	}
}
