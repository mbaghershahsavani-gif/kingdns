package routing

func ExecuteRuntime(domain string) Decision {
	return Decision{
		Domain: domain,
		Target: "127.0.0.1",
		Reason: "controller",
	}
}
