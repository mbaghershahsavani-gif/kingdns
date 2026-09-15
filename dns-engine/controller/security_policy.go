package controller

type SecurityPolicy struct {
	Name    string
	Enabled bool
}

func ApplySecurityPolicy(name string) SecurityPolicy {
	return SecurityPolicy{
		Name:    name,
		Enabled: true,
	}
}
