package controller

type Node struct {
	ID      int
	Name    string
	Health  int
	Latency int
	Region  string
}

func Available() []Node {
	return []Node{}
}
