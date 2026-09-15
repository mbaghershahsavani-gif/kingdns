package controller

type CloudProvider struct {
	Name      string
	Connected bool
}

func ConnectProvider(name string) CloudProvider {
	return CloudProvider{
		Name:      name,
		Connected: true,
	}
}
