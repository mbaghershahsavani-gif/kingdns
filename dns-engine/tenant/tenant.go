package tenant

type Tenant struct {
	ID     string
	Name   string
	Active bool
}

func Create(name string) Tenant {
	return Tenant{
		Name:   name,
		Active: true,
	}
}
