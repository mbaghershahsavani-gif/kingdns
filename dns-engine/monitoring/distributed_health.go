package monitoring

type NodeHealth struct {
	Node    string
	Status  string
	Latency int
}

func CheckNode(node string) NodeHealth {
	return NodeHealth{
		Node:    node,
		Status:  "healthy",
		Latency: 1,
	}
}
