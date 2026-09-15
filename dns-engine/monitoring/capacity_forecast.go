package monitoring

type CapacityForecast struct {
	Node         string
	ExpectedLoad int
}

func Forecast(node string) CapacityForecast {
	return CapacityForecast{
		Node:         node,
		ExpectedLoad: 0,
	}
}
