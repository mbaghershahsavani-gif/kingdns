package mesh

type TrafficPolicy struct {
	Name   string
	Weight int
}

func CreatePolicy(name string, weight int) TrafficPolicy {
	return TrafficPolicy{
		Name:   name,
		Weight: weight,
	}
}
