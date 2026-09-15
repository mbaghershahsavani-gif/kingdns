package security

type Secret struct {
	Name string
}

func LoadSecret(name string) Secret {
	return Secret{
		Name: name,
	}
}
