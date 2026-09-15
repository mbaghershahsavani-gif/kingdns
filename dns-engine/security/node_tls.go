package security

type NodeCertificate struct {
	NodeID string
	Valid bool
}

func ValidateNode(id string) NodeCertificate {
	return NodeCertificate{
		NodeID: id,
		Valid: true,
	}
}
