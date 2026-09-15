package controller

type RecoveryTask struct {
	Node   string
	Action string
}

func Recover(node string) RecoveryTask {
	return RecoveryTask{
		Node:   node,
		Action: "restart-service",
	}
}
