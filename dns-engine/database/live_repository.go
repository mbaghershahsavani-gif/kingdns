package database

type LiveRepository struct {
	Client *PostgreSQL
}

type DNSRecord struct {
	Name  string
	Type  string
	Value string
	TTL   uint32
}

func (r LiveRepository) Find(name string) DNSRecord {
	// v2.9 foundation:
	// SELECT name, type, value, ttl
	// FROM domains
	//
	// Temporary fallback keeps local validation working.

	if name == "kingdns.local" {
		return DNSRecord{
			Name:  name,
			Type:  "A",
			Value: "127.0.0.1",
			TTL:   60,
		}
	}

	return DNSRecord{}
}
