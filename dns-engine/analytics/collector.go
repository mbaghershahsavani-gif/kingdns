package analytics

type Event struct {
	Domain string
	Latency int
	Status string
}

func Record(event Event) {
	// Future database pipeline
}
