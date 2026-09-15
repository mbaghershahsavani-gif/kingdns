package auth

type Role string

const (
	Admin Role = "ADMIN"
	Operator Role = "OPERATOR"
	NodeAgent Role = "NODE_AGENT"
	Viewer Role = "VIEWER"
)
