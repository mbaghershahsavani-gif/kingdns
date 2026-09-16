package cluster

func Bootstrap() {

	UpdateHeartbeat()

	RegisterScore(
		CalculateScore(),
	)
}
