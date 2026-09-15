package cache

type RedisClient struct {
}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (r *RedisClient) Get(domain string) string {
	return ""
}

func (r *RedisClient) Set(domain string, value string, ttl int) {
}
