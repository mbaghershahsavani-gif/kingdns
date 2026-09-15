package security

type Identity struct {
	Name  string
	Valid bool
}

func Authenticate(name string) Identity {
	return Identity{
		Name:  name,
		Valid: true,
	}
}
