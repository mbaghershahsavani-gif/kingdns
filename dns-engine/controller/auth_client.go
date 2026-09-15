package controller

type AuthClient struct {
	Token string
}

func NewAuthClient(token string) *AuthClient {
	return &AuthClient{Token: token}
}
