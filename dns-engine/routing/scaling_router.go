package routing

type ScalingDecision struct {
	Scale  bool
	Target string
}

func DecideScaling(load int) ScalingDecision {
	return ScalingDecision{
		Scale:  load > 1000,
		Target: "default-node",
	}
}
