package routing

type Policy struct {
	Mode   string
	Region string
}

func Apply(policy Policy) string {
	return policy.Mode
}
