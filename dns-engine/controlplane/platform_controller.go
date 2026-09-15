package controlplane

type PlatformController struct {
	Name   string
	Active bool
}

func Start(name string) PlatformController {
	return PlatformController{
		Name:   name,
		Active: true,
	}
}
