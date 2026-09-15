package routing

type RecoveryNode struct {
	IP        string
	Available bool
}

func SelectRecovery(nodes []RecoveryNode) string {
	for _, node := range nodes {
		if node.Available {
			return node.IP
		}
	}

	return ""
}
