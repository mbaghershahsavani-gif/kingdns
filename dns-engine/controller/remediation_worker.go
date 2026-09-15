package controller

type RemediationTask struct {
	Node   string
	Action string
}

func ExecuteRemediation(node string) RemediationTask {
	return RemediationTask{
		Node:   node,
		Action: "restore-service",
	}
}
