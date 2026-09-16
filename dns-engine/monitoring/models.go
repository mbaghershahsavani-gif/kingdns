package monitoring

type Record struct {
	Name  string
	Value string
}

type DNSQueryLog struct {
	Domain   string
	Response string
}

type AnalyticsEvent struct {
	Name  string
	Value string
}

type NodeHealth struct {
	Node    string
	Status  string
	Healthy bool
	Latency int
}
