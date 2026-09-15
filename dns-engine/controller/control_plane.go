package controller

type ControlPlane struct {
	Token string
}

func NewControlPlane(token string) *ControlPlane {
	return &ControlPlane{Token: token}
}

func (c *ControlPlane) Sync() bool {
	return c.Token != ""
}
