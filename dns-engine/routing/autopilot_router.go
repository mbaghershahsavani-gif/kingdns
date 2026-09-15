package routing

type AutopilotNode struct {
	IP    string
	Score int
}

func ChooseBest(nodes []AutopilotNode) string {
	best := ""
	score := -1

	for _, node := range nodes {
		if node.Score > score {
			score = node.Score
			best = node.IP
		}
	}

	return best
}
