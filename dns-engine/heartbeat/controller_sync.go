package heartbeat

type NodeStatus struct {
	NodeID int
	Status string
}

func Sync(status NodeStatus) {
}
