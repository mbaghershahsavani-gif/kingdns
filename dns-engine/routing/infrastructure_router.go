package routing

type InfrastructureNode struct {
	IP       string
	Capacity int
	Healthy  bool
}

func SelectInfrastructureNode(nodes []InfrastructureNode) string {
	for _, node := range nodes {
		if node.Healthy {
			return node.IP
		}
	}
	return ""
}
