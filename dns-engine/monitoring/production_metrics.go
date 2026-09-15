package monitoring

type Metric struct {
	Name  string
	Value int
}

func Collect() []Metric {
	return []Metric{}
}
