package intelligence

type OptimizationDecision struct {
	Target string
	Score  int
}

func Optimize(target string) OptimizationDecision {
	return OptimizationDecision{
		Target: target,
		Score:  0,
	}
}
