package controller

type OperationMetric struct {
	Name  string
	Value int
}

func GetMetrics() []OperationMetric {
	return []OperationMetric{}
}
