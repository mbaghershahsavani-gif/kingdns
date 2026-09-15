package controller

type ConsensusState struct {
	Leader string
	Nodes  int
}

func ElectLeader(nodes int) ConsensusState {
	return ConsensusState{
		Leader: "kingdns-controller",
		Nodes:  nodes,
	}
}
