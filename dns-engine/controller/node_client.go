package controller

type Node struct {
	ID int
	Name string
	Health int
	Latency int
}

func HealthyNodes() []Node {
	return []Node{}
}
