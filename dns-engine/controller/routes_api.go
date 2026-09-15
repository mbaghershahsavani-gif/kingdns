package controller

type APIRoute struct {
	Domain   string
	NodeID   int
	Priority int
}

func (c *Client) FetchRoutes() ([]APIRoute, error) {
	return []APIRoute{}, nil
}
