package resolver

func ARecord(domain string, ip string) DNSRecord {
	return DNSRecord{
		Name: domain,
		Type: "A",
		Value: ip,
		TTL: 60,
	}
}
