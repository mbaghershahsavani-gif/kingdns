package security

type AgentIdentity struct {
    NodeID int
    Token string
}

func Validate(identity AgentIdentity) bool {
    return identity.Token != ""
}
