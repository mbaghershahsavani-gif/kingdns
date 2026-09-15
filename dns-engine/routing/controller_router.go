package routing

type Node struct {
	ID      int
	Name    string
	IP      string
	Healthy bool
	Latency int
}

func SelectBestNode(nodes []Node) Node {
	for _, node := range nodes {
		if node.Healthy {
			return node
		}
	}

	return Node{}
}
