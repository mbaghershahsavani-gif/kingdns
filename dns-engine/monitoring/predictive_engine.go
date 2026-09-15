package monitoring

type PredictionResult struct {
	Node        string
	Probability int
}

func PredictFailure(node string) PredictionResult {
	return PredictionResult{
		Node:        node,
		Probability: 0,
	}
}
