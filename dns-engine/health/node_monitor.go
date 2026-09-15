package health

type Node struct {
	Name string
	Score int
}

func Available(node Node) bool {
	return node.Score > 0
}
