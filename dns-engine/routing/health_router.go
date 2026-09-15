package routing

type Node struct {
	Name string
	Health int
}

func SelectHealthy(nodes []Node) Node {
	for _, node := range nodes {
		if node.Health > 0 {
			return node
		}
	}

	return Node{}
}
