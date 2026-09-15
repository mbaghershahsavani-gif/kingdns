package routing

type Decision struct {
	Node   string
	Reason string
}

func Resolve() Decision {
	return Decision{
		Reason: "health-latency policy",
	}
}
