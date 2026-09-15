package monitoring

func CheckNode(node string) NodeHealth {
	return NodeHealth{
		Node:    node,
		Status:  "healthy",
		Latency: 1,
	}
}
