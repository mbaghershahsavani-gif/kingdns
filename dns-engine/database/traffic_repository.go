package database

type TrafficRecord struct {
	Region  string
	Queries int
	Latency int
}

func GetTrafficStats() []TrafficRecord {
	return []TrafficRecord{
		{
			Region:  "local",
			Queries: 1,
			Latency: 1,
		},
	}
}
