package monitoring

type SystemState struct {
	Node    string
	Healthy bool
}

func Evaluate(state SystemState) string {
	if state.Healthy {
		return "stable"
	}

	return "remediation-required"
}
