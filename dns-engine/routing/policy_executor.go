package routing

func ExecutePolicy(policy Policy) Decision {
	return Decision{
		Reason: "policy-execution",
	}
}
