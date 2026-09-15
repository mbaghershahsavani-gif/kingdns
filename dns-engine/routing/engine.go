package routing

func ResolveLegacy() Decision {
	return Decision{
		Reason: "health-latency policy",
	}
}
