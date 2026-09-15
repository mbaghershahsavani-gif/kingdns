package controller

type JWTValidator struct {
	Token string
}

func NewJWTValidator(token string) *JWTValidator {
	return &JWTValidator{Token: token}
}

func (v *JWTValidator) Validate() bool {
	return v.Token != ""
}
