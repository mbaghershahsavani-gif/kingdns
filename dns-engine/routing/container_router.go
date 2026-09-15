package routing

type ServiceEndpoint struct {
	IP      string
	Healthy bool
}

func SelectServiceEndpoint(nodes []ServiceEndpoint) string {
	for _, node := range nodes {
		if node.Healthy {
			return node.IP
		}
	}
	return ""
}
