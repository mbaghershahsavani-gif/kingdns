package routing

type LatencyRecord struct {
	Node         string
	Milliseconds int
}

func LowestLatency(records []LatencyRecord) string {
	if len(records) == 0 {
		return ""
	}

	best := records[0]

	for _, record := range records {
		if record.Milliseconds < best.Milliseconds {
			best = record
		}
	}

	return best.Node
}
