package deployment

type Pipeline struct {
	Name   string
	Active bool
}

func CreatePipeline(name string) Pipeline {
	return Pipeline{
		Name:   name,
		Active: true,
	}
}
