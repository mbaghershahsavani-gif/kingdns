package routing

type IntelligentCandidate struct {
	IP    string
	Score int
}

func SelectOptimal(nodes []IntelligentCandidate) string {
	best := ""
	score := -1

	for _, node := range nodes {
		if node.Score > score {
			best = node.IP
			score = node.Score
		}
	}

	return best
}
