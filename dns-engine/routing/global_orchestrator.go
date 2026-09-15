package routing

type GlobalRoute struct {
	Region string
	Score  int
}

func SelectGlobalRoute(routes []GlobalRoute) string {
	best := ""
	score := -1

	for _, route := range routes {
		if route.Score > score {
			best = route.Region
			score = route.Score
		}
	}

	return best
}
