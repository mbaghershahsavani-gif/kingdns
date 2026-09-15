package controller

type AuthenticatedClient struct {
	Token string
}

func NewAuthenticatedClient(token string) *AuthenticatedClient {
	return &AuthenticatedClient{Token: token}
}

func (c *AuthenticatedClient) HealthyNodes() []string {
	return []string{"127.0.0.1"}
}
