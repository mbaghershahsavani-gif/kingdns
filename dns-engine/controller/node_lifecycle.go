package controller

type NodeStatus string

const (
	Registering NodeStatus = "REGISTERING"
	Active      NodeStatus = "ACTIVE"
	Degraded    NodeStatus = "DEGRADED"
	Recovering  NodeStatus = "RECOVERING"
	Failed      NodeStatus = "FAILED"
	Restored    NodeStatus = "RESTORED"
)

type NodeLifecycle struct {
	Node   string
	Status NodeStatus
}

func UpdateNode(node string, status NodeStatus) NodeLifecycle {
	return NodeLifecycle{
		Node:   node,
		Status: status,
	}
}
