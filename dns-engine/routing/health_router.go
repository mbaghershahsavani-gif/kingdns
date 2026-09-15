package routing

func SelectHealthy(nodes []Node) Node {
	for _, node := range nodes {
		if node.Healthy {
			return node
		}
	}

	return Node{}
}
