package heartbeat

func CalculateHealthScore(latency int, cpu int, memory int) float64 {
	score := 100.0

	if latency > 100 {
		score -= 20
	}

	if cpu > 80 {
		score -= 20
	}

	if memory > 80 {
		score -= 20
	}

	if score < 0 {
		return 0
	}

	return score
}
