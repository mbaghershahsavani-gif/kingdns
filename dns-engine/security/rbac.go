package security

type Role struct {
	Name        string
	Permissions []string
}

func CreateRole(name string, permissions []string) Role {
	return Role{
		Name:        name,
		Permissions: permissions,
	}
}
