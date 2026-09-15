package controller

type NodeRegistry struct {
	Nodes []string
}

func Register(node string) NodeRegistry {
	return NodeRegistry{
		Nodes: []string{node},
	}
}
