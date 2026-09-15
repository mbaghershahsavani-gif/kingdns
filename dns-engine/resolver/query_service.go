package resolver

type QueryService struct {
}

func (q QueryService) Resolve(domain string) DNSRecord {
	return BuildAnswer(domain, "127.0.0.1")
}
