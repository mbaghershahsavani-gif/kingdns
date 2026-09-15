package routing

type NodeScore struct {
	NodeID uint
	Latency int
	Health float64
}

// Future:
// SelectBestNode will combine latency, health and geo rules.
func SelectBestNode(nodes []NodeScore) uint {
	if len(nodes) == 0 {
		return 0
	}
	return nodes[0].NodeID
}
