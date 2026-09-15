package cache

type DNSCache struct {
}

func (c DNSCache) Get(domain string) string {
	return ""
}

func (c DNSCache) Set(domain string, value string) {
}
