package routing

type Policy struct {
	Mode string
}

func ExecutePolicy(policy Policy) Decision {
	return Decision{}
}
