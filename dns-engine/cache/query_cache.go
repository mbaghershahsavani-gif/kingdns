package cache

type QueryCache struct {
}

func (c QueryCache) Get(domain string) string {
	return ""
}

func (c QueryCache) Set(domain string, value string, ttl uint32) {
}
