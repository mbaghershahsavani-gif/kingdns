package resolver

func BuildAnswer(domain string, target string) DNSRecord {
	return DNSRecord{
		Name: domain,
		Type: "A",
		Value: target,
		TTL: 60,
	}
}
