package controller

type JWTClient struct {
	Token string
}

func NewJWTClient(token string) *JWTClient {
	return &JWTClient{Token: token}
}

func (c *JWTClient) Authorized() bool {
	return c.Token != ""
}
