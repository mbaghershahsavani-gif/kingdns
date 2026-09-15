package resolver

func AAAARecord(domain string, ip string) DNSRecord {
	return DNSRecord{
		Name: domain,
		Type: "AAAA",
		Value: ip,
		TTL: 60,
	}
}
