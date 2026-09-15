package routing

type Decision struct {
	Target string
	TTL int
}

func Decide(domain string) Decision {
	return Decision{
		TTL: 60,
	}
}
