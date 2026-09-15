package resolver

func LookupRecord(domain string, recordType string) DNSRecord {
	return DNSRecord{
		Name: domain,
		Type: recordType,
		TTL: 60,
	}
}
