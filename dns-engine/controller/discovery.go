package controller

type Node struct {
	Name    string
	Health  int
	Latency int
}

func Discover() []Node {
	return []Node{}
}
