package monitoring

type HealthScore struct {
	Availability int
	Latency      int
	Errors       int
	Score        int
}

func CalculateScore(h HealthScore) int {
	return h.Score
}
