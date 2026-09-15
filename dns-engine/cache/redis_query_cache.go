package cache

type RedisQueryCache struct {
}

func (c RedisQueryCache) Get(domain string) string {
	return ""
}

func (c RedisQueryCache) Set(domain string, value string, ttl int) {
}
