package monitoring

type HealthState struct {
	Node  string
	Score int
}

func SyncHealth(node string, score int) HealthState {
	return HealthState{
		Node:  node,
		Score: score,
	}
}
