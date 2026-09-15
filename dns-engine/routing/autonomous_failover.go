package routing

type AutonomousNode struct {
	IP    string
	Score int
}

func SelectAutonomous(nodes []AutonomousNode) string {
	best := ""

	for _, node := range nodes {
		if node.Score > 0 {
			best = node.IP
			break
		}
	}

	return best
}
