package controller

func (c *Client) GetAuthenticatedNodes() ([]Node, error) {
	// v3.2 foundation:
	// Authorization: Bearer <token>
	// GET /api/nodes

	return []Node{
		{
			Name:    "local-test-node",
			Health:  100,
			Latency: 1,
			Region:  "IR",
		},
	}, nil
}
