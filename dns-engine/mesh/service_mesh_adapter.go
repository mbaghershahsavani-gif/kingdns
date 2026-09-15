package mesh

type ServiceEndpoint struct {
	Name    string
	Healthy bool
}

func RegisterService(name string) ServiceEndpoint {
	return ServiceEndpoint{
		Name:    name,
		Healthy: true,
	}
}
