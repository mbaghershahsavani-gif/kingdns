package cache

type RedisService struct {
}

func (r RedisService) Get(domain string) string {
	return ""
}

func (r RedisService) Set(domain string, value string, ttl int) {
}
