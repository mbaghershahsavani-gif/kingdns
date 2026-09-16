package cluster

type Decision struct {
	ActiveNode   string `json:"active_node"`
	ActiveRegion string `json:"active_region"`
	State        string `json:"state"`
	Score        int    `json:"score"`
	Reason       string `json:"reason"`
}

func CalculateDecision() Decision {

	nodes := GetNodes()

	preferred := PreferredRegion()

	best := Decision{
		Score:  -1,
		State:  string(Offline),
		Reason: "no healthy nodes available",
	}

	// First try preferred region

	for _, node := range nodes {

		state := EvaluateHealth(node)

		if node.Region == preferred &&
			state != Offline {

			return Decision{
				ActiveNode:   node.NodeID,
				ActiveRegion: node.Region,
				State:        string(state),
				Score:        node.Total,
				Reason:       "preferred region healthy",
			}
		}
	}

	// Fallback to highest healthy node

	for _, node := range nodes {

		state := EvaluateHealth(node)

		if state != Offline &&
			node.Total > best.Score {

			best.ActiveNode = node.NodeID
			best.ActiveRegion = node.Region
			best.State = string(state)
			best.Score = node.Total
			best.Reason = "fallback highest health score"
		}
	}

	return best
}
