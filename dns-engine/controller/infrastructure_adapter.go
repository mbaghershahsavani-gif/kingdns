package controller

type InfrastructureAction struct {
	Node   string
	Action string
}

func ExecuteInfrastructureAction(node string, action string) InfrastructureAction {
	return InfrastructureAction{
		Node:   node,
		Action: action,
	}
}
