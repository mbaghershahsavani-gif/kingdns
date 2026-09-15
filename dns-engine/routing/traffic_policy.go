package routing

type TrafficPolicy struct {
	Name   string
	Weight int
	Region string
}

func ApplyTrafficPolicy(policies []TrafficPolicy) string {
	for _, policy := range policies {
		if policy.Weight > 0 {
			return policy.Region
		}
	}

	return ""
}
