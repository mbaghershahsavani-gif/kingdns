package cluster

import (
	"time"
)

type Decision struct {
	ActiveNode   string `json:"active_node"`
	ActiveRegion string `json:"active_region"`
	State        string `json:"state"`
	Score        int    `json:"score"`
	Reason       string `json:"reason"`
}

const nodeTimeout = 90 * time.Second

func IsNodeAlive(node Score) bool {

	return time.Since(node.LastSeen) < nodeTimeout
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

		if !IsNodeAlive(node) {
			continue
		}

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

	// Fallback to highest healthy live node

	for _, node := range nodes {

		if !IsNodeAlive(node) {
			continue
		}

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
