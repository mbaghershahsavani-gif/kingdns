package resolver

type RecordRepository struct {
}

func (r RecordRepository) Find(domain string) DNSRecord {
	return DNSRecord{
		Name: domain,
		Type: "A",
		TTL: 60,
	}
}
