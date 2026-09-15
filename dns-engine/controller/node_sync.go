package controller

type NodeState struct {
	Name   string
	Health int
}

func SyncNode(state NodeState) NodeState {
	return state
}
