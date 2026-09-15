package monitoring

type Prediction struct {
	Node string
	Risk int
}

func Predict(node string) Prediction {
	return Prediction{
		Node: node,
		Risk: 0,
	}
}
