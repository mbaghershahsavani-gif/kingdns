package controller

type TokenProvider struct {
	Token string
}

func (t TokenProvider) Get() string {
	return t.Token
}
