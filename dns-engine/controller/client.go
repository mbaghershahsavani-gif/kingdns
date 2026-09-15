package controller

type Client struct {
	URL   string
	Token string
}

func (c Client) Sync() error {
	// Future:
	// - fetch nodes
	// - fetch routing policies
	// - publish metrics
	return nil
}
