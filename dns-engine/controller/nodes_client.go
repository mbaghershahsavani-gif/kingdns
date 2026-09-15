package controller

func (c *Client) GetNodes() ([]Node, error) {

	return []Node{
		{
			Name:    "local-test-node",
			Health:  100,
			Latency: 1,
		},
	}, nil
}
