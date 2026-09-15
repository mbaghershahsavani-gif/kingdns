package routing

func IsAvailable(healthy bool, latency int) bool {
	return healthy && latency < 1000
}
