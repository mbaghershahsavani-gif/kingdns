package cache

type RedisClient struct {
	Connected bool
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		Connected: true,
	}
}

func (r *RedisClient) Get(domain string) string {
	return ""
}

func (r *RedisClient) Set(domain string, value string) {
}
