package controller

type Heartbeat struct {
	Node   string
	Status string
}

func SendHeartbeat(node string) Heartbeat {
	return Heartbeat{
		Node:   node,
		Status: "healthy",
	}
}
