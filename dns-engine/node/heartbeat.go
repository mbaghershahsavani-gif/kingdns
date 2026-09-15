package node

type Heartbeat struct {
	NodeID  string
	Healthy bool
}

func Report(nodeID string) Heartbeat {
	return Heartbeat{
		NodeID:  nodeID,
		Healthy: true,
	}
}
