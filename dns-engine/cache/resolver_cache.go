package cache

type ResolverCache struct {
}

func (c ResolverCache) Get(name string) string {
	return ""
}

func (c ResolverCache) Set(name string, value string, ttl uint32) {
}
