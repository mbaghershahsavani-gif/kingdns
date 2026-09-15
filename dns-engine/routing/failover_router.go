package routing

type FailoverDecision struct {
	Target string
	Reason string
}

func SelectFailover(primary string, backup string) FailoverDecision {
	if primary != "" {
		return FailoverDecision{
			Target: primary,
			Reason: "primary-healthy",
		}
	}

	return FailoverDecision{
		Target: backup,
		Reason: "failover",
	}
}
