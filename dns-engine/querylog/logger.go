package querylog

type Entry struct {
	Domain  string
	Status  string
	Latency int
}

func Write(entry Entry) {
	// Future:
	// PostgreSQL analytics pipeline
}
