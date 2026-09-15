package controller

type APINode struct {
	ID        int
	Name      string
	Type      string
	Country   string
	IPAddress string
	Status    string
	Health    int
	Latency   int
}

func (c *Client) FetchNodes() ([]APINode, error) {
	// v3.3 foundation.
	// Next step:
	// real GET /api/nodes with JWT authentication.

	return []APINode{
		{
			ID:        1,
			Name:      "local-test-node",
			Type:      "relay",
			Country:   "IR",
			IPAddress: "127.0.0.1",
			Status:    "online",
			Health:    100,
			Latency:   1,
		},
	}, nil
}
