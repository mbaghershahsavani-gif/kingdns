package cluster

func Bootstrap() {

	UpdateHeartbeat()

	score := CalculateScore()

	RegisterScore(score)
}
