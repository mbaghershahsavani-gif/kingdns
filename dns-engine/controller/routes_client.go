package controller

type Route struct {
	Domain string
	NodeID int
}

func (c *Client) GetRoutes() ([]Route, error) {
	// Future:
	// GET /api/routes

	return []Route{}, nil
}
