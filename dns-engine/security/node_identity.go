package security

type NodeIdentity struct {
	Node     string
	Verified bool
}

func VerifyNode(node string) NodeIdentity {
	return NodeIdentity{
		Node:     node,
		Verified: true,
	}
}
