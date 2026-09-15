package routing

type Decision struct {
	Domain string
	Target string
	Reason string
}

func Decide(domain string) Decision {
	return Decision{
		Domain: domain,
		Target: "127.0.0.1",
		Reason: "default-routing",
	}
}
