package resolver

func CNAMERecord(domain string, target string) DNSRecord {
	return DNSRecord{
		Name:  domain,
		Type:  "CNAME",
		Value: target,
		TTL:   60,
	}
}
