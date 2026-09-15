package controller

type CloudResource struct {
	Name   string
	Status string
}

func Provision(name string) CloudResource {
	return CloudResource{
		Name:   name,
		Status: "ready",
	}
}
