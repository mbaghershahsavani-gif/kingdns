package monitoring

type DNSQueryLog struct {
	Domain  string
	Node    string
	Latency int
	Status  string
}

func Record(log DNSQueryLog) {
}
