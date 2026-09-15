package controller

type APIClient struct {
	BaseURL string
	Token   string
}

func (c APIClient) GetNodes() error {
	// Future:
	// - call controller API
	// - retrieve active DNS nodes
	return nil
}
