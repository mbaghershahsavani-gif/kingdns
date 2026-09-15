package resolver

type PlatformResolver struct {
}

func (p PlatformResolver) Resolve(domain string) DNSRecord {
	return DNSRecord{
		Name: domain,
		Type: "A",
		TTL:  60,
	}
}
