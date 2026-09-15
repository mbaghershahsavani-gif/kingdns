package routing

type MeshNode struct {
	Region string
	IP     string
	Score  int
}

func SelectMeshNode(nodes []MeshNode) string {
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
