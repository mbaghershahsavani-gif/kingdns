package monitoring

type MeshHealth struct {
	Service string
	Healthy bool
}

func CheckMesh(service string) MeshHealth {
	return MeshHealth{
		Service: service,
		Healthy: true,
	}
}
