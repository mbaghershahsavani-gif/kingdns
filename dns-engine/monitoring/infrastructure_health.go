package monitoring

type InfrastructureHealth struct {
	Node      string
	Available bool
}

func CheckInfrastructure(node string) InfrastructureHealth {
	return InfrastructureHealth{
		Node:      node,
		Available: true,
	}
}
